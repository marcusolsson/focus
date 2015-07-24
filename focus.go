package focus

import (
	"bytes"
	"log"
	"time"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

// Event ...
type Event interface {
}

// Gained ...
type Gained struct {
	WMClass []string
}

// Lost ...
type Lost struct {
	WMClass []string
}

var msRefresh time.Duration = 200

// Listen ...
func Listen(fn func(e Event)) error {
	X, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	setup := xproto.Setup(X)
	root := setup.DefaultScreen(X).Root

	var prevWindowID xproto.Window
	var prevClasses []string

	for {
		<-time.After(msRefresh * time.Millisecond)

		// From one of the xgb examples.
		aname := "_NET_ACTIVE_WINDOW"
		activeAtom, err := xproto.InternAtom(X, true, uint16(len(aname)),
			aname).Reply()
		if err != nil {
			log.Println(err)
			continue
		}

		reply, err := xproto.GetProperty(X, false, root, activeAtom.Atom,
			xproto.GetPropertyTypeAny, 0, (1<<32)-1).Reply()
		if err != nil {
			log.Println(err)
			continue
		}

		windowID := xproto.Window(xgb.Get32(reply.Value))

		aname = "WM_CLASS"
		nameAtom, err := xproto.InternAtom(X, true, uint16(len(aname)),
			aname).Reply()
		if err != nil {
			log.Println(err)
			continue
		}

		reply, err = xproto.GetProperty(X, false, windowID, nameAtom.Atom,
			xproto.GetPropertyTypeAny, 0, (1<<32)-1).Reply()
		if err != nil {
			log.Println(err)
			continue
		}

		classes := stringSlice(reply.Value)

		// Check if active window has changed since last time.
		if windowID != prevWindowID {
			fn(Lost{
				WMClass: prevClasses,
			})
			fn(Gained{
				WMClass: classes,
			})
			prevWindowID = windowID
			prevClasses = classes
		}
	}
}

func stringSlice(b []byte) []string {
	// Strings are separated by zeros.
	bs := bytes.Split(b, []byte{0})

	var result []string
	for _, v := range bs {
		if len(v) > 0 {
			result = append(result, string(v))
		}
	}

	return result
}
