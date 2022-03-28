package krptool

import "encoding/xml"

type Krpano struct {
	XMLName        xml.Name `xml:"krpano"`
	Version        string   `xml:"version,attr"`
	Onstart        string   `xml:"onstart,attr"`
	Basedir        string   `xml:"basedir,attr"`
	Bgcolor        string   `xml:"bgcolor,attr"`
	Idletime       float64  `xml:"idletime,attr"`
	Strict         bool     `xml:"strict,attr"`
	Showerrors     bool     `xml:"showerrors,attr"`
	Logkey         bool     `xml:"logkey,attr"`
	Debugmode      bool     `xml:"debugmode,attr"`
	Debugkeys      bool     `xml:"debugkeys,attr"`
	Debugjsactions bool     `xml:"debugjsactions,attr"`
	Debugjsplugins bool     `xml:"debugjsplugins,attr"`
	// sub tags
	Include     []Include   `xml:"include"`
	Preview     Preview     `xml:"preview"`
	Image       Image       `xml:"image"`
	View        View        `xml:"view"`
	Area        Area        `xml:"area"`
	Display     Display     `xml:"display"`
	Control     Control     `xml:"control"`
	Cursors     Cursors     `xml:"cursors"`
	Autorotate  Autorotate  `xml:"autorotate"`
	Plugin      []Plugin    `xml:"plugin"`
	Layer       []Layer     `xml:"layer"`
	Hotspot     []Hotspot   `xml:"hotspot"`
	Events      []Events    `xml:"events"`
	Action      []Action    `xml:"action"`
	Contextmenu Contextmenu `xml:"contextmenu"`
	Network     Network     `xml:"network"`
	Memory      Memory      `xml:"memory"`
	Security    Security    `xml:"security"`
	Textstyle   Textstyle   `xml:"textstyle"`
	Data        []Data      `xml:"data"`
	Scene       []Scene     `xml:"scene"`
	Set         []Set       `xml:"set"`
}

type Include struct {
	XMLName xml.Name `xml:"include"`
	Url     string   `xml:"url,attr"`
}

type Preview struct {
	XMLName    xml.Name `xml:"preview"`
	Type       string   `xml:"type,attr"`
	Url        string   `xml:"url,attr"`
	Striporder string   `xml:"striporder,attr"`
}

type Image struct {
	XMLName           xml.Name `xml:"image"`
	Type              string   `xml:"type,attr"`
	Hfov              float64  `xml:"hfov,attr"`
	Vfov              float64  `xml:"vfov,attr"`
	Voffset           float64  `xml:"voffset,attr"`
	Crop              string   `xml:"crop,attr"`
	Prealign          string   `xml:"prealign,attr"`
	Multiresthreshold float64  `xml:"multiresthreshold,attr"`
	Baseindex         int      `xml:"baseindex,attr"`
	Tilesize          int      `xml:"tilesize,attr"`
	Tileoverlap       int      `xml:"tileoverlap,attr"`
	Cubelabels        string   `xml:"cubelabels,attr"`
	Cubemapmode       bool     `xml:"cubemapmode,attr"`
	Cubemapsize       int      `xml:"cubemapsize,attr"`
	Stereo            bool     `xml:"stereo,attr"`
	Stereolabels      string   `xml:"stereolabels,attr"`
	Stereoformat      string   `xml:"stereoformat,attr"`
	Stereooffset      float64  `xml:"stereooffset,attr"`
	Mjpegstream       string   `xml:"mjpegstream,attr"`
	Ox                float64  `xml:"ox,attr"`
	Oy                float64  `xml:"oy,attr"`
	Oz                float64  `xml:"oz,attr"`
	Frames            int      `xml:"frames,attr"`
	Frame             int      `xml:"frame,attr"`
	Loadstate         int      `xml:"loadstate,attr"`
	// sub tags
	Depthmap Depthmap `xml:"depthmap"`
}

type Depthmap struct {
	XMLName     xml.Name `xml:"depthmap"`
	Url         string   `xml:"url,attr"`
	Enabled     bool     `xml:"enabled,attr"`
	Rendermode  string   `xml:"rendermode,attr"`
	Background  string   `xml:"background,attr"`
	Scale       float64  `xml:"scale,attr"`
	Offset      float64  `xml:"offset,attr"`
	Subdiv      float64  `xml:"subdiv,attr"`
	Encoding    string   `xml:"encoding,attr"`
	Axis        string   `xml:"axis,attr"`
	Cull        string   `xml:"cull,attr"`
	Center      string   `xml:"center,attr"`
	Waitforload bool     `xml:"waitforload,attr"`
}

type View struct {
	XMLName           xml.Name `xml:"view"`
	Hlookat           float64  `xml:"hlookat,attr"`
	Vlookat           float64  `xml:"vlookat,attr"`
	Camroll           float64  `xml:"camroll,attr"`
	Fovtype           string   `xml:"fovtype,attr"`
	Fov               float64  `xml:"fov,attr"`
	Hfov              float64  `xml:"hfov,attr"`
	Vfov              float64  `xml:"vfov,attr"`
	Dfov              float64  `xml:"dfov,attr"`
	Mfov              float64  `xml:"mfov,attr"`
	Sfov              float64  `xml:"sfov,attr"`
	Fovmin            float64  `xml:"fovmin,attr"`
	Fovmax            float64  `xml:"fovmax,attr"`
	Maxpixelzoom      float64  `xml:"maxpixelzoom,attr"`
	Mfovratio         float64  `xml:"mfovratio,attr"`
	Distortion        float64  `xml:"distortion,attr"`
	Fisheye           float64  `xml:"fisheye,attr"`
	Distortionfovlink float64  `xml:"distortionfovlink,attr"`
	Fisheyefovlink    float64  `xml:"fisheyefovlink,attr"`
	Stereographic     bool     `xml:"stereographic,attr"`
	Pannini           float64  `xml:"pannini,attr"`
	Architectural     float64  `xml:"architectural,attr"`
	Limitview         string   `xml:"limitview,attr"`
	Hlookatmin        float64  `xml:"hlookatmin,attr"`
	Hlookatmax        float64  `xml:"hlookatmax,attr"`
	Vlookatmin        float64  `xml:"vlookatmin,attr"`
	Vlookatmax        float64  `xml:"vlookatmax,attr"`
	Hlookatrange      float64  `xml:"hlookatrange,attr"`
	Vlookatrange      float64  `xml:"vlookatrange,attr"`
	Rx                float64  `xml:"rx,attr"`
	Ry                float64  `xml:"ry,attr"`
	Tx                float64  `xml:"tx,attr"`
	Ty                float64  `xml:"ty,attr"`
	Tz                float64  `xml:"tz,attr"`
	Ox                float64  `xml:"ox,attr"`
	Oy                float64  `xml:"oy,attr"`
	Oz                float64  `xml:"oz,attr"`
}

type Area struct {
	XMLName    xml.Name `xml:"area"`
	Mode       string   `xml:"mode,attr"`
	Align      string   `xml:"align,attr"`
	X          string   `xml:"x,attr"`
	Y          string   `xml:"y,attr"`
	Width      string   `xml:"width,attr"`
	Height     string   `xml:"height,attr"`
	Left       string   `xml:"left,attr"`
	Top        string   `xml:"top,attr"`
	Right      string   `xml:"right,attr"`
	Bottom     string   `xml:"bottom,attr"`
	Cliplayers bool     `xml:"cliplayers,attr"`
}

type Display struct {
	XMLName              xml.Name `xml:"display"`
	Autofullscreen       bool     `xml:"autofullscreen,attr"`
	Stereo               bool     `xml:"stereo,attr"`
	Stereooverlap        float64  `xml:"stereooverlap,attr"`
	Stereoscale          float64  `xml:"stereoscale,attr"`
	Monoside             int      `xml:"monoside,attr"`
	Anaglyph             bool     `xml:"anaglyph,attr"`
	Safearea             string   `xml:"safearea,attr"`
	SafeareaInset        string   `xml:"safearea_inset,attr"`
	Mipmapping           string   `xml:"mipmapping,attr"`
	Loadwhilemoving      string   `xml:"loadwhilemoving,attr"`
	Framebufferscale     float64  `xml:"framebufferscale,attr"`
	Hotspotvrrendering   string   `xml:"hotspotvrrendering,attr"`
	Hotspotworldscale    float64  `xml:"hotspotworldscale,attr"`
	Depthmaprendermode   string   `xml:"depthmaprendermode,attr"`
	Depthbuffer          bool     `xml:"depthbuffer,attr"`
	Depthrange           string   `xml:"depthrange,attr"`
	Surfacesubdiv        int      `xml:"surfacesubdiv,attr"`
	Nofullscreenfallback bool     `xml:"nofullscreenfallback,attr"`
	Iphonefullscreen     int      `xml:"iphonefullscreen,attr"`
	Showpolys            bool     `xml:"showpolys,attr"`
}

type Control struct {
	XMLName          xml.Name `xml:"control"`
	Usercontrol      string   `xml:"usercontrol,attr"`
	Mode             string   `xml:"mode,attr"`
	Dragrelative     bool     `xml:"dragrelative,attr"`
	Draginertia      float64  `xml:"draginertia,attr"`
	Dragfriction     float64  `xml:"dragfriction,attr"`
	DragOldmode      bool     `xml:"drag_oldmode,attr"`
	Movetorelative   bool     `xml:"movetorelative,attr"`
	Movetoaccelerate float64  `xml:"movetoaccelerate,attr"`
	Movetospeed      float64  `xml:"movetospeed,attr"`
	Movetofriction   float64  `xml:"movetofriction,attr"`
	Movetoyfriction  float64  `xml:"movetoyfriction,attr"`
	Frictionstop     float64  `xml:"frictionstop,attr"`
	Invert           bool     `xml:"invert,attr"`
	Keybaccelerate   float64  `xml:"keybaccelerate,attr"`
	Keybspeed        float64  `xml:"keybspeed,attr"`
	Keybfriction     float64  `xml:"keybfriction,attr"`
	Keybinvert       bool     `xml:"keybinvert,attr"`
	Keybfovchange    float64  `xml:"keybfovchange,attr"`
	Mousefovchange   float64  `xml:"mousefovchange,attr"`
	Fovspeed         float64  `xml:"fovspeed,attr"`
	Fovfriction      float64  `xml:"fovfriction,attr"`
	Zoomtocursor     bool     `xml:"zoomtocursor,attr"`
	Zoomoutcursor    bool     `xml:"zoomoutcursor,attr"`
	Disablewheel     bool     `xml:"disablewheel,attr"`
	Touchzoom        bool     `xml:"touchzoom,attr"`
	Keycodesleft     string   `xml:"keycodesleft,attr"`
	Keycodesright    string   `xml:"keycodesright,attr"`
	Keycodesup       string   `xml:"keycodesup,attr"`
	Keycodesdown     string   `xml:"keycodesdown,attr"`
	Keycodesin       string   `xml:"keycodesin,attr"`
	Keycodesout      string   `xml:"keycodesout,attr"`
	Keydownrepeat    bool     `xml:"keydownrepeat,attr"`
	Bouncinglimits   string   `xml:"bouncinglimits,attr"`
}

type Cursors struct {
	XMLName  xml.Name `xml:"cursors"`
	Standard string   `xml:"standard,attr"`
	Dragging string   `xml:"dragging,attr"`
	Moving   string   `xml:"moving,attr"`
	Url      string   `xml:"url,attr"`
	Type     string   `xml:"type,attr"`
	Move     string   `xml:"move,attr"`
	Drag     string   `xml:"drag,attr"`
	ArrowL   string   `xml:"arrow_l,attr"`
	ArrowR   string   `xml:"arrow_r,attr"`
	ArrowU   string   `xml:"arrow_u,attr"`
	ArrowD   string   `xml:"arrow_d,attr"`
	ArrowLu  string   `xml:"arrow_lu,attr"`
	ArrowRu  string   `xml:"arrow_ru,attr"`
	ArrowLd  string   `xml:"arrow_ld,attr"`
	ArrowRd  string   `xml:"arrow_rd,attr"`
}

type Autorotate struct {
	XMLName            xml.Name `xml:"autorotate"`
	Enabled            bool     `xml:"enabled,attr"`
	Waittime           float64  `xml:"waittime,attr"`
	Accel              float64  `xml:"accel,attr"`
	Speed              float64  `xml:"speed,attr"`
	Horizon            float64  `xml:"horizon,attr"`
	Tofov              float64  `xml:"tofov,attr"`
	Oneroundrange      float64  `xml:"oneroundrange,attr"`
	Zoomslowdown       bool     `xml:"zoomslowdown,attr"`
	Interruptionevents string   `xml:"interruptionevents,attr"`
	Start              string   `xml:"start,attr"`
	Stop               string   `xml:"stop,attr"`
	Interrupt          string   `xml:"interrupt,attr"`
	Pause              string   `xml:"pause,attr"`
	Resume             string   `xml:"resume,attr"`
}

type Plugin struct {
	XMLName       xml.Name `xml:"plugin"`
	Name          string   `xml:"name,attr"`
	Type          string   `xml:"type,attr"`
	Url           string   `xml:"url,attr"`
	Keep          bool     `xml:"keep,attr"`
	Visible       bool     `xml:"visible,attr"`
	Enabled       bool     `xml:"enabled,attr"`
	Handcursor    bool     `xml:"handcursor,attr"`
	Cursor        string   `xml:"cursor,attr"`
	Maskchildren  bool     `xml:"maskchildren,attr"`
	Zorder        string   `xml:"zorder,attr"`
	Align         string   `xml:"align,attr"`
	Edge          string   `xml:"edge,attr"`
	Safearea      bool     `xml:"safearea,attr"`
	X             string   `xml:"x,attr"`
	Y             string   `xml:"y,attr"`
	Rotate        float64  `xml:"rotate,attr"`
	Width         string   `xml:"width,attr"`
	Height        string   `xml:"height,attr"`
	Scale         float64  `xml:"scale,attr"`
	Alpha         float64  `xml:"alpha,attr"`
	Crop          string   `xml:"crop,attr"`
	Onovercrop    string   `xml:"onovercrop,attr"`
	Ondowncrop    string   `xml:"ondowncrop,attr"`
	Parent        string   `xml:"parent,attr"`
	Scalechildren bool     `xml:"scalechildren,attr"`
	Bgcolor       int      `xml:"bgcolor,attr"`
	Bgalpha       float64  `xml:"bgalpha,attr"`
	Bgborder      string   `xml:"bgborder,attr"`
	Bgroundedge   string   `xml:"bgroundedge,attr"`
	Bgshadow      string   `xml:"bgshadow,attr"`
	Bgcapture     bool     `xml:"bgcapture,attr"`
	Onover        string   `xml:"onover,attr"`
	Onhover       string   `xml:"onhover,attr"`
	Onout         string   `xml:"onout,attr"`
	Onclick       string   `xml:"onclick,attr"`
	Ondown        string   `xml:"ondown,attr"`
	Onup          string   `xml:"onup,attr"`
	Onloaded      string   `xml:"onloaded,attr"`
}

type Layer struct {
	Plugin
	Layer *[]Layer `xml:"layer"`
}

type Hotspot struct {
	XMLName           xml.Name `xml:"hotspot"`
	Enabled           bool     `xml:"enabled,attr"`
	Visible           bool     `xml:"visible,attr"`
	Alpha             float64  `xml:"alpha,attr"`
	Layer             string   `xml:"layer,attr"`
	Name              string   `xml:"name,attr"`
	Index             int      `xml:"index,attr"`
	Type              string   `xml:"type,attr"`
	Url               string   `xml:"url,attr"`
	Keep              bool     `xml:"keep,attr"`
	Renderer          string   `xml:"renderer,attr"`
	Handcursor        bool     `xml:"handcursor,attr"`
	Cursor            string   `xml:"cursor,attr"`
	Maskchildren      bool     `xml:"maskchildren,attr"`
	Zorder            string   `xml:"zorder,attr"`
	Capture           bool     `xml:"capture,attr"`
	Capturefocus      bool     `xml:"capturefocus,attr"`
	Children          bool     `xml:"children,attr"`
	Nativecontextmenu bool     `xml:"nativecontextmenu,attr"`
	Blendmode         string   `xml:"blendmode,attr"`
	Alphachannel      string   `xml:"alphachannel,attr"`
	Chromakey         string   `xml:"chromakey,attr"`
	Mipmapping        bool     `xml:"mipmapping,attr"`
	Premultiplyalpha  bool     `xml:"premultiplyalpha,attr"`
	Ath               float64  `xml:"ath,attr"`
	Atv               float64  `xml:"atv,attr"`
	Edge              string   `xml:"edge,attr"`
	Ox                string   `xml:"ox,attr"`
	Oy                string   `xml:"oy,attr"`
	Oref              int      `xml:"oref,attr"`
	Zoom              bool     `xml:"zoom,attr"`
	Distorted         bool     `xml:"distorted,attr"`
	Rotationorder     string   `xml:"rotationorder,attr"`
	Rx                float64  `xml:"rx,attr"`
	Ry                float64  `xml:"ry,attr"`
	Rz                float64  `xml:"rz,attr"`
	Inverserotation   bool     `xml:"inverserotation,attr"`
	Depth             float64  `xml:"depth,attr"`
	Depthbuffer       bool     `xml:"depthbuffer,attr"`
	Torigin           string   `xml:"torigin,attr"`
	Tx                float64  `xml:"tx,attr"`
	Ty                float64  `xml:"ty,attr"`
	Tz                float64  `xml:"tz,attr"`
	Prealign          bool     `xml:"prealign,attr"`
	Flying            float64  `xml:"flying,attr"`
	Scaleflying       bool     `xml:"scaleflying,attr"`
	Width             string   `xml:"width,attr"`
	Height            string   `xml:"height,attr"`
	Imagewidth        int      `xml:"imagewidth,attr"`
	Imageheight       int      `xml:"imageheight,attr"`
	Scale             float64  `xml:"scale,attr"`
	Rotate            float64  `xml:"rotate,attr"`
	Pixelhittest      bool     `xml:"pixelhittest,attr"`
	Accuracy          int      `xml:"accuracy,attr"`
	Autoalpha         bool     `xml:"autoalpha,attr"`
	Usecontentsize    bool     `xml:"usecontentsize,attr"`
	Scale9Grid        string   `xml:"scale9grid,attr"`
	Stereo            string   `xml:"stereo,attr"`
	Crop              string   `xml:"crop,attr"`
	Onovercrop        string   `xml:"onovercrop,attr"`
	Ondowncrop        string   `xml:"ondowncrop,attr"`
	Scalechildren     bool     `xml:"scalechildren,attr"`
	Polyline          bool     `xml:"polyline,attr"`
	Fillcolor         int      `xml:"fillcolor,attr"`
	Fillalpha         float64  `xml:"fillalpha,attr"`
	Borderwidth       float64  `xml:"borderwidth,attr"`
	Bordercolor       int      `xml:"bordercolor,attr"`
	Borderalpha       float64  `xml:"borderalpha,attr"`
	Borderzoom        float64  `xml:"borderzoom,attr"`
	Borderhittest     bool     `xml:"borderhittest,attr"`
	Subdiv            bool     `xml:"subdiv,attr"`
	Ishotspot         bool     `xml:"ishotspot,attr"`
	Loading           bool     `xml:"loading,attr"`
	Loaded            bool     `xml:"loaded,attr"`
	Loadedurl         string   `xml:"loadedurl,attr"`
	Hovering          bool     `xml:"hovering,attr"`
	Pressed           bool     `xml:"pressed,attr"`
	Hitx              float64  `xml:"hitx,attr"`
	Hity              float64  `xml:"hity,attr"`
	Hitd              float64  `xml:"hitd,attr"`
	Onover            string   `xml:"onover,attr"`
	Onhover           string   `xml:"onhover,attr"`
	Onout             string   `xml:"onout,attr"`
	Ondown            string   `xml:"ondown,attr"`
	Onup              string   `xml:"onup,attr"`
	Onclick           string   `xml:"onclick,attr"`
	Onloaded          string   `xml:"onloaded,attr"`
}

type Events struct {
	XMLName              xml.Name `xml:"events"`
	Name                 string   `xml:"name,attr"`
	Keep                 bool     `xml:"keep,attr"`
	Onenterfullscreen    string   `xml:"onenterfullscreen,attr"`
	Onexitfullscreen     string   `xml:"onexitfullscreen,attr"`
	Onxmlcomplete        string   `xml:"onxmlcomplete,attr"`
	Onpreviewcomplete    string   `xml:"onpreviewcomplete,attr"`
	Onloadcomplete       string   `xml:"onloadcomplete,attr"`
	Onblendcomplete      string   `xml:"onblendcomplete,attr"`
	Onnewpano            string   `xml:"onnewpano,attr"`
	Onremovepano         string   `xml:"onremovepano,attr"`
	Onnewscene           string   `xml:"onnewscene,attr"`
	Onxmlerror           string   `xml:"onxmlerror,attr"`
	Onloaderror          string   `xml:"onloaderror,attr"`
	Onkeydown            string   `xml:"onkeydown,attr"`
	Onkeyup              string   `xml:"onkeyup,attr"`
	Onclick              string   `xml:"onclick,attr"`
	Onsingleclick        string   `xml:"onsingleclick,attr"`
	Ondoubleclick        string   `xml:"ondoubleclick,attr"`
	Onmousedown          string   `xml:"onmousedown,attr"`
	Onmouseup            string   `xml:"onmouseup,attr"`
	Onmousewheel         string   `xml:"onmousewheel,attr"`
	Oncontextmenu        string   `xml:"oncontextmenu,attr"`
	Onidle               string   `xml:"onidle,attr"`
	Onviewchange         string   `xml:"onviewchange,attr"`
	Onviewchanged        string   `xml:"onviewchanged,attr"`
	Onresize             string   `xml:"onresize,attr"`
	Onframebufferresize  string   `xml:"onframebufferresize,attr"`
	Onautorotatestart    string   `xml:"onautorotatestart,attr"`
	Onautorotatestop     string   `xml:"onautorotatestop,attr"`
	Onautorotateoneround string   `xml:"onautorotateoneround,attr"`
	Onautorotatechange   string   `xml:"onautorotatechange,attr"`
	Oniphonefullscreen   string   `xml:"oniphonefullscreen,attr"`
}

type Action struct {
	XMLName xml.Name `xml:"action"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
	Scope   string   `xml:"scope,attr"`
	Args    string   `xml:"args,attr"`
	Autorun string   `xml:"autorun,attr"`
	Protect bool     `xml:"protect,attr"`
	Secure  bool     `xml:"secure,attr"`
	Content string   `xml:"content,attr"`
}

type Contextmenu struct {
	XMLName     xml.Name `xml:"contextmenu"`
	Fullscreen  bool     `xml:"fullscreen,attr"`
	Versioninfo bool     `xml:"versioninfo,attr"`
	Touch       bool     `xml:"touch,attr"`
	Customstyle string   `xml:"customstyle,attr"`
	Enterfs     string   `xml:"enterfs,attr"`
	Exitfs      string   `xml:"exitfs,attr"`
}

type Network struct {
	XMLName    xml.Name `xml:"network"`
	Retrycount int      `xml:"retrycount,attr"`
}

type Memory struct {
	XMLName xml.Name `xml:"memory"`
	Maxmem  int      `xml:"maxmem,attr"`
}

type Security struct {
	XMLName xml.Name `xml:"security"`
	Cors    string   `xml:"cors,attr"`
}

type Textstyle struct {
	XMLName xml.Name `xml:"textstyle"`
}

type Data struct {
	XMLName xml.Name `xml:"data"`
	Name    int      `xml:"name,attr"`
	Content string   `xml:"content,attr"`
}

type Scene struct {
	XMLName  xml.Name `xml:"scene"`
	Name     string   `xml:"name,attr"`
	Index    int      `xml:"index,attr"`
	Onstart  string   `xml:"onstart,attr"`
	Autoload bool     `xml:"autoload,attr"`
	Content  string   `xml:"content,attr"`

	// sub tags
	Include    []Include  `xml:"include"`
	Preview    Preview    `xml:"preview"`
	Image      Image      `xml:"image"`
	View       View       `xml:"view"`
	Area       Area       `xml:"area"`
	Display    Display    `xml:"display"`
	Control    Control    `xml:"control"`
	Cursors    Cursors    `xml:"cursors"`
	Autorotate Autorotate `xml:"autorotate"`
	Plugin     []Plugin   `xml:"plugin"`
	Hotspot    []Hotspot  `xml:"hotspot"`
	Events     []Events   `xml:"events"`
	Action     []Action   `xml:"action"`
	Data       []Data     `xml:"data"`
	Set        []Set      `xml:"set"`
}

type Set struct {
	XMLName xml.Name `xml:"set"`
	Var     string   `xml:"var,attr"`
	Val     string   `xml:"val,attr"`
}
