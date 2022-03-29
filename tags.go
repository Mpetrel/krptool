package krptool

import "encoding/xml"

type Krpano struct {
	XMLName        xml.Name `xml:"krpano"`
	Version        string   `xml:"version,attr,omitempty"`
	Onstart        string   `xml:"onstart,attr,omitempty"`
	Basedir        string   `xml:"basedir,attr,omitempty"`
	Bgcolor        string   `xml:"bgcolor,attr,omitempty"`
	Idletime       float64  `xml:"idletime,attr,omitempty"`
	Strict         bool     `xml:"strict,attr,omitempty"`
	Showerrors     bool     `xml:"showerrors,attr,omitempty"`
	Logkey         bool     `xml:"logkey,attr,omitempty"`
	Debugmode      bool     `xml:"debugmode,attr,omitempty"`
	Debugkeys      bool     `xml:"debugkeys,attr,omitempty"`
	Debugjsactions bool     `xml:"debugjsactions,attr,omitempty"`
	Debugjsplugins bool     `xml:"debugjsplugins,attr,omitempty"`
	// sub tags
	Include     []*Include   `xml:"include"`
	Preview     *Preview     `xml:"preview"`
	Image       *Image       `xml:"image"`
	View        *View        `xml:"view"`
	Area        *Area        `xml:"area"`
	Display     *Display     `xml:"display"`
	Control     *Control     `xml:"control"`
	Cursors     *Cursors     `xml:"cursors"`
	Autorotate  *Autorotate  `xml:"autorotate"`
	Plugin      []*Plugin    `xml:"plugin"`
	Layer       []*Layer     `xml:"layer"`
	Hotspot     []*Hotspot   `xml:"hotspot"`
	Events      []*Events    `xml:"events"`
	Action      []*Action    `xml:"action"`
	Contextmenu *Contextmenu `xml:"contextmenu"`
	Network     *Network     `xml:"network"`
	Memory      *Memory      `xml:"memory"`
	Security    *Security    `xml:"security"`
	Textstyle   *Textstyle   `xml:"textstyle"`
	Data        []*Data      `xml:"data"`
	Scene       []*Scene     `xml:"scene"`
	Set         []*Set       `xml:"set"`
}

type Include struct {
	XMLName xml.Name `xml:"include"`
	Url     string   `xml:"url,attr,omitempty"`
}

type Preview struct {
	XMLName    xml.Name `xml:"preview"`
	Type       string   `xml:"type,attr,omitempty"`
	Url        string   `xml:"url,attr,omitempty"`
	Striporder string   `xml:"striporder,attr,omitempty"`
}

type Image struct {
	XMLName           xml.Name `xml:"image"`
	Style             string   `xml:"style,attr,omitempty"`
	Type              string   `xml:"type,attr,omitempty"`
	Hfov              float64  `xml:"hfov,attr,omitempty"`
	Vfov              float64  `xml:"vfov,attr,omitempty"`
	Voffset           float64  `xml:"voffset,attr,omitempty"`
	Crop              string   `xml:"crop,attr,omitempty"`
	Prealign          string   `xml:"prealign,attr,omitempty"`
	Multiresthreshold float64  `xml:"multiresthreshold,attr,omitempty"`
	Baseindex         int      `xml:"baseindex,attr,omitempty"`
	Tilesize          int      `xml:"tilesize,attr,omitempty"`
	Tileoverlap       int      `xml:"tileoverlap,attr,omitempty"`
	Cubelabels        string   `xml:"cubelabels,attr,omitempty"`
	Cubemapmode       bool     `xml:"cubemapmode,attr,omitempty"`
	Cubemapsize       int      `xml:"cubemapsize,attr,omitempty"`
	Stereo            bool     `xml:"stereo,attr,omitempty"`
	Stereolabels      string   `xml:"stereolabels,attr,omitempty"`
	Stereoformat      string   `xml:"stereoformat,attr,omitempty"`
	Stereooffset      float64  `xml:"stereooffset,attr,omitempty"`
	Mjpegstream       string   `xml:"mjpegstream,attr,omitempty"`
	Ox                float64  `xml:"ox,attr,omitempty"`
	Oy                float64  `xml:"oy,attr,omitempty"`
	Oz                float64  `xml:"oz,attr,omitempty"`
	Frames            int      `xml:"frames,attr,omitempty"`
	Frame             int      `xml:"frame,attr,omitempty"`
	Loadstate         int      `xml:"loadstate,attr,omitempty"`
	// sub tags
	Cube     *Cube     `xml:"cube"`
	Depthmap *Depthmap `xml:"depthmap"`
}

type Cube struct {
	XMLName  xml.Name `xml:"cube"`
	Multires string   `xml:"multires,attr,omitempty"`
	Url      string   `xml:"url,attr,omitempty"`
}

type Depthmap struct {
	XMLName     xml.Name `xml:"depthmap"`
	Url         string   `xml:"url,attr,omitempty"`
	Enabled     bool     `xml:"enabled,attr,omitempty"`
	Rendermode  string   `xml:"rendermode,attr,omitempty"`
	Background  string   `xml:"background,attr,omitempty"`
	Scale       float64  `xml:"scale,attr,omitempty"`
	Offset      float64  `xml:"offset,attr,omitempty"`
	Subdiv      float64  `xml:"subdiv,attr,omitempty"`
	Encoding    string   `xml:"encoding,attr,omitempty"`
	Axis        string   `xml:"axis,attr,omitempty"`
	Cull        string   `xml:"cull,attr,omitempty"`
	Center      string   `xml:"center,attr,omitempty"`
	Waitforload bool     `xml:"waitforload,attr,omitempty"`
}

type View struct {
	XMLName           xml.Name `xml:"view"`
	Hlookat           float64  `xml:"hlookat,attr,omitempty"`
	Vlookat           float64  `xml:"vlookat,attr,omitempty"`
	Camroll           float64  `xml:"camroll,attr,omitempty"`
	Fovtype           string   `xml:"fovtype,attr,omitempty"`
	Fov               float64  `xml:"fov,attr,omitempty"`
	Hfov              float64  `xml:"hfov,attr,omitempty"`
	Vfov              float64  `xml:"vfov,attr,omitempty"`
	Dfov              float64  `xml:"dfov,attr,omitempty"`
	Mfov              float64  `xml:"mfov,attr,omitempty"`
	Sfov              float64  `xml:"sfov,attr,omitempty"`
	Fovmin            float64  `xml:"fovmin,attr,omitempty"`
	Fovmax            float64  `xml:"fovmax,attr,omitempty"`
	Maxpixelzoom      float64  `xml:"maxpixelzoom,attr,omitempty"`
	Mfovratio         float64  `xml:"mfovratio,attr,omitempty"`
	Distortion        float64  `xml:"distortion,attr,omitempty"`
	Fisheye           float64  `xml:"fisheye,attr,omitempty"`
	Distortionfovlink float64  `xml:"distortionfovlink,attr,omitempty"`
	Fisheyefovlink    float64  `xml:"fisheyefovlink,attr,omitempty"`
	Stereographic     bool     `xml:"stereographic,attr,omitempty"`
	Pannini           float64  `xml:"pannini,attr,omitempty"`
	Architectural     float64  `xml:"architectural,attr,omitempty"`
	Limitview         string   `xml:"limitview,attr,omitempty"`
	Hlookatmin        float64  `xml:"hlookatmin,attr,omitempty"`
	Hlookatmax        float64  `xml:"hlookatmax,attr,omitempty"`
	Vlookatmin        float64  `xml:"vlookatmin,attr,omitempty"`
	Vlookatmax        float64  `xml:"vlookatmax,attr,omitempty"`
	Hlookatrange      float64  `xml:"hlookatrange,attr,omitempty"`
	Vlookatrange      float64  `xml:"vlookatrange,attr,omitempty"`
	Rx                float64  `xml:"rx,attr,omitempty"`
	Ry                float64  `xml:"ry,attr,omitempty"`
	Tx                float64  `xml:"tx,attr,omitempty"`
	Ty                float64  `xml:"ty,attr,omitempty"`
	Tz                float64  `xml:"tz,attr,omitempty"`
	Ox                float64  `xml:"ox,attr,omitempty"`
	Oy                float64  `xml:"oy,attr,omitempty"`
	Oz                float64  `xml:"oz,attr,omitempty"`
}

type Area struct {
	XMLName    xml.Name `xml:"area"`
	Mode       string   `xml:"mode,attr,omitempty"`
	Align      string   `xml:"align,attr,omitempty"`
	X          string   `xml:"x,attr,omitempty"`
	Y          string   `xml:"y,attr,omitempty"`
	Width      string   `xml:"width,attr,omitempty"`
	Height     string   `xml:"height,attr,omitempty"`
	Left       string   `xml:"left,attr,omitempty"`
	Top        string   `xml:"top,attr,omitempty"`
	Right      string   `xml:"right,attr,omitempty"`
	Bottom     string   `xml:"bottom,attr,omitempty"`
	Cliplayers bool     `xml:"cliplayers,attr,omitempty"`
}

type Display struct {
	XMLName              xml.Name `xml:"display"`
	Autofullscreen       bool     `xml:"autofullscreen,attr,omitempty"`
	Stereo               bool     `xml:"stereo,attr,omitempty"`
	Stereooverlap        float64  `xml:"stereooverlap,attr,omitempty"`
	Stereoscale          float64  `xml:"stereoscale,attr,omitempty"`
	Monoside             int      `xml:"monoside,attr,omitempty"`
	Anaglyph             bool     `xml:"anaglyph,attr,omitempty"`
	Safearea             string   `xml:"safearea,attr,omitempty"`
	SafeareaInset        string   `xml:"safearea_inset,attr,omitempty"`
	Mipmapping           string   `xml:"mipmapping,attr,omitempty"`
	Loadwhilemoving      string   `xml:"loadwhilemoving,attr,omitempty"`
	Framebufferscale     float64  `xml:"framebufferscale,attr,omitempty"`
	Hotspotvrrendering   string   `xml:"hotspotvrrendering,attr,omitempty"`
	Hotspotworldscale    float64  `xml:"hotspotworldscale,attr,omitempty"`
	Depthmaprendermode   string   `xml:"depthmaprendermode,attr,omitempty"`
	Depthbuffer          bool     `xml:"depthbuffer,attr,omitempty"`
	Depthrange           string   `xml:"depthrange,attr,omitempty"`
	Surfacesubdiv        int      `xml:"surfacesubdiv,attr,omitempty"`
	Nofullscreenfallback bool     `xml:"nofullscreenfallback,attr,omitempty"`
	Iphonefullscreen     int      `xml:"iphonefullscreen,attr,omitempty"`
	Showpolys            bool     `xml:"showpolys,attr,omitempty"`
}

type Control struct {
	XMLName          xml.Name `xml:"control"`
	Usercontrol      string   `xml:"usercontrol,attr,omitempty"`
	Mode             string   `xml:"mode,attr,omitempty"`
	Dragrelative     bool     `xml:"dragrelative,attr,omitempty"`
	Draginertia      float64  `xml:"draginertia,attr,omitempty"`
	Dragfriction     float64  `xml:"dragfriction,attr,omitempty"`
	DragOldmode      bool     `xml:"drag_oldmode,attr,omitempty"`
	Movetorelative   bool     `xml:"movetorelative,attr,omitempty"`
	Movetoaccelerate float64  `xml:"movetoaccelerate,attr,omitempty"`
	Movetospeed      float64  `xml:"movetospeed,attr,omitempty"`
	Movetofriction   float64  `xml:"movetofriction,attr,omitempty"`
	Movetoyfriction  float64  `xml:"movetoyfriction,attr,omitempty"`
	Frictionstop     float64  `xml:"frictionstop,attr,omitempty"`
	Invert           bool     `xml:"invert,attr,omitempty"`
	Keybaccelerate   float64  `xml:"keybaccelerate,attr,omitempty"`
	Keybspeed        float64  `xml:"keybspeed,attr,omitempty"`
	Keybfriction     float64  `xml:"keybfriction,attr,omitempty"`
	Keybinvert       bool     `xml:"keybinvert,attr,omitempty"`
	Keybfovchange    float64  `xml:"keybfovchange,attr,omitempty"`
	Mousefovchange   float64  `xml:"mousefovchange,attr,omitempty"`
	Fovspeed         float64  `xml:"fovspeed,attr,omitempty"`
	Fovfriction      float64  `xml:"fovfriction,attr,omitempty"`
	Zoomtocursor     bool     `xml:"zoomtocursor,attr,omitempty"`
	Zoomoutcursor    bool     `xml:"zoomoutcursor,attr,omitempty"`
	Disablewheel     bool     `xml:"disablewheel,attr,omitempty"`
	Touchzoom        bool     `xml:"touchzoom,attr,omitempty"`
	Keycodesleft     string   `xml:"keycodesleft,attr,omitempty"`
	Keycodesright    string   `xml:"keycodesright,attr,omitempty"`
	Keycodesup       string   `xml:"keycodesup,attr,omitempty"`
	Keycodesdown     string   `xml:"keycodesdown,attr,omitempty"`
	Keycodesin       string   `xml:"keycodesin,attr,omitempty"`
	Keycodesout      string   `xml:"keycodesout,attr,omitempty"`
	Keydownrepeat    bool     `xml:"keydownrepeat,attr,omitempty"`
	Bouncinglimits   string   `xml:"bouncinglimits,attr,omitempty"`
}

type Cursors struct {
	XMLName  xml.Name `xml:"cursors"`
	Standard string   `xml:"standard,attr,omitempty"`
	Dragging string   `xml:"dragging,attr,omitempty"`
	Moving   string   `xml:"moving,attr,omitempty"`
	Url      string   `xml:"url,attr,omitempty"`
	Type     string   `xml:"type,attr,omitempty"`
	Move     string   `xml:"move,attr,omitempty"`
	Drag     string   `xml:"drag,attr,omitempty"`
	ArrowL   string   `xml:"arrow_l,attr,omitempty"`
	ArrowR   string   `xml:"arrow_r,attr,omitempty"`
	ArrowU   string   `xml:"arrow_u,attr,omitempty"`
	ArrowD   string   `xml:"arrow_d,attr,omitempty"`
	ArrowLu  string   `xml:"arrow_lu,attr,omitempty"`
	ArrowRu  string   `xml:"arrow_ru,attr,omitempty"`
	ArrowLd  string   `xml:"arrow_ld,attr,omitempty"`
	ArrowRd  string   `xml:"arrow_rd,attr,omitempty"`
}

type Autorotate struct {
	XMLName            xml.Name `xml:"autorotate"`
	Enabled            bool     `xml:"enabled,attr,omitempty"`
	Waittime           float64  `xml:"waittime,attr,omitempty"`
	Accel              float64  `xml:"accel,attr,omitempty"`
	Speed              float64  `xml:"speed,attr,omitempty"`
	Horizon            float64  `xml:"horizon,attr,omitempty"`
	Tofov              float64  `xml:"tofov,attr,omitempty"`
	Oneroundrange      float64  `xml:"oneroundrange,attr,omitempty"`
	Zoomslowdown       bool     `xml:"zoomslowdown,attr,omitempty"`
	Interruptionevents string   `xml:"interruptionevents,attr,omitempty"`
	Start              string   `xml:"start,attr,omitempty"`
	Stop               string   `xml:"stop,attr,omitempty"`
	Interrupt          string   `xml:"interrupt,attr,omitempty"`
	Pause              string   `xml:"pause,attr,omitempty"`
	Resume             string   `xml:"resume,attr,omitempty"`
}

type Plugin struct {
	XMLName       xml.Name `xml:"plugin"`
	Name          string   `xml:"name,attr,omitempty"`
	Type          string   `xml:"type,attr,omitempty"`
	Url           string   `xml:"url,attr,omitempty"`
	Keep          bool     `xml:"keep,attr,omitempty"`
	Visible       bool     `xml:"visible,attr,omitempty"`
	Enabled       bool     `xml:"enabled,attr,omitempty"`
	Handcursor    bool     `xml:"handcursor,attr,omitempty"`
	Cursor        string   `xml:"cursor,attr,omitempty"`
	Maskchildren  bool     `xml:"maskchildren,attr,omitempty"`
	Zorder        string   `xml:"zorder,attr,omitempty"`
	Align         string   `xml:"align,attr,omitempty"`
	Edge          string   `xml:"edge,attr,omitempty"`
	Safearea      bool     `xml:"safearea,attr,omitempty"`
	X             string   `xml:"x,attr,omitempty"`
	Y             string   `xml:"y,attr,omitempty"`
	Rotate        float64  `xml:"rotate,attr,omitempty"`
	Width         string   `xml:"width,attr,omitempty"`
	Height        string   `xml:"height,attr,omitempty"`
	Scale         float64  `xml:"scale,attr,omitempty"`
	Alpha         float64  `xml:"alpha,attr,omitempty"`
	Crop          string   `xml:"crop,attr,omitempty"`
	Onovercrop    string   `xml:"onovercrop,attr,omitempty"`
	Ondowncrop    string   `xml:"ondowncrop,attr,omitempty"`
	Parent        string   `xml:"parent,attr,omitempty"`
	Scalechildren bool     `xml:"scalechildren,attr,omitempty"`
	Bgcolor       int      `xml:"bgcolor,attr,omitempty"`
	Bgalpha       float64  `xml:"bgalpha,attr,omitempty"`
	Bgborder      string   `xml:"bgborder,attr,omitempty"`
	Bgroundedge   string   `xml:"bgroundedge,attr,omitempty"`
	Bgshadow      string   `xml:"bgshadow,attr,omitempty"`
	Bgcapture     bool     `xml:"bgcapture,attr,omitempty"`
	Onover        string   `xml:"onover,attr,omitempty"`
	Onhover       string   `xml:"onhover,attr,omitempty"`
	Onout         string   `xml:"onout,attr,omitempty"`
	Onclick       string   `xml:"onclick,attr,omitempty"`
	Ondown        string   `xml:"ondown,attr,omitempty"`
	Onup          string   `xml:"onup,attr,omitempty"`
	Onloaded      string   `xml:"onloaded,attr,omitempty"`
}

type Layer struct {
	Plugin
	Layer *[]Layer `xml:"layer"`
}

type Hotspot struct {
	XMLName           xml.Name `xml:"hotspot"`
	Enabled           bool     `xml:"enabled,attr,omitempty"`
	Visible           bool     `xml:"visible,attr,omitempty"`
	Alpha             float64  `xml:"alpha,attr,omitempty"`
	Layer             string   `xml:"layer,attr,omitempty"`
	Name              string   `xml:"name,attr,omitempty"`
	Index             int      `xml:"index,attr,omitempty"`
	Type              string   `xml:"type,attr,omitempty"`
	Url               string   `xml:"url,attr,omitempty"`
	Keep              bool     `xml:"keep,attr,omitempty"`
	Renderer          string   `xml:"renderer,attr,omitempty"`
	Handcursor        bool     `xml:"handcursor,attr,omitempty"`
	Cursor            string   `xml:"cursor,attr,omitempty"`
	Maskchildren      bool     `xml:"maskchildren,attr,omitempty"`
	Zorder            string   `xml:"zorder,attr,omitempty"`
	Capture           bool     `xml:"capture,attr,omitempty"`
	Capturefocus      bool     `xml:"capturefocus,attr,omitempty"`
	Children          bool     `xml:"children,attr,omitempty"`
	Nativecontextmenu bool     `xml:"nativecontextmenu,attr,omitempty"`
	Blendmode         string   `xml:"blendmode,attr,omitempty"`
	Alphachannel      string   `xml:"alphachannel,attr,omitempty"`
	Chromakey         string   `xml:"chromakey,attr,omitempty"`
	Mipmapping        bool     `xml:"mipmapping,attr,omitempty"`
	Premultiplyalpha  bool     `xml:"premultiplyalpha,attr,omitempty"`
	Ath               float64  `xml:"ath,attr,omitempty"`
	Atv               float64  `xml:"atv,attr,omitempty"`
	Edge              string   `xml:"edge,attr,omitempty"`
	Ox                string   `xml:"ox,attr,omitempty"`
	Oy                string   `xml:"oy,attr,omitempty"`
	Oref              int      `xml:"oref,attr,omitempty"`
	Zoom              bool     `xml:"zoom,attr,omitempty"`
	Distorted         bool     `xml:"distorted,attr,omitempty"`
	Rotationorder     string   `xml:"rotationorder,attr,omitempty"`
	Rx                float64  `xml:"rx,attr,omitempty"`
	Ry                float64  `xml:"ry,attr,omitempty"`
	Rz                float64  `xml:"rz,attr,omitempty"`
	Inverserotation   bool     `xml:"inverserotation,attr,omitempty"`
	Depth             float64  `xml:"depth,attr,omitempty"`
	Depthbuffer       bool     `xml:"depthbuffer,attr,omitempty"`
	Torigin           string   `xml:"torigin,attr,omitempty"`
	Tx                float64  `xml:"tx,attr,omitempty"`
	Ty                float64  `xml:"ty,attr,omitempty"`
	Tz                float64  `xml:"tz,attr,omitempty"`
	Prealign          bool     `xml:"prealign,attr,omitempty"`
	Flying            float64  `xml:"flying,attr,omitempty"`
	Scaleflying       bool     `xml:"scaleflying,attr,omitempty"`
	Width             string   `xml:"width,attr,omitempty"`
	Height            string   `xml:"height,attr,omitempty"`
	Imagewidth        int      `xml:"imagewidth,attr,omitempty"`
	Imageheight       int      `xml:"imageheight,attr,omitempty"`
	Scale             float64  `xml:"scale,attr,omitempty"`
	Rotate            float64  `xml:"rotate,attr,omitempty"`
	Pixelhittest      bool     `xml:"pixelhittest,attr,omitempty"`
	Accuracy          int      `xml:"accuracy,attr,omitempty"`
	Autoalpha         bool     `xml:"autoalpha,attr,omitempty"`
	Usecontentsize    bool     `xml:"usecontentsize,attr,omitempty"`
	Scale9Grid        string   `xml:"scale9grid,attr,omitempty"`
	Stereo            string   `xml:"stereo,attr,omitempty"`
	Crop              string   `xml:"crop,attr,omitempty"`
	Onovercrop        string   `xml:"onovercrop,attr,omitempty"`
	Ondowncrop        string   `xml:"ondowncrop,attr,omitempty"`
	Scalechildren     bool     `xml:"scalechildren,attr,omitempty"`
	Polyline          bool     `xml:"polyline,attr,omitempty"`
	Fillcolor         int      `xml:"fillcolor,attr,omitempty"`
	Fillalpha         float64  `xml:"fillalpha,attr,omitempty"`
	Borderwidth       float64  `xml:"borderwidth,attr,omitempty"`
	Bordercolor       int      `xml:"bordercolor,attr,omitempty"`
	Borderalpha       float64  `xml:"borderalpha,attr,omitempty"`
	Borderzoom        float64  `xml:"borderzoom,attr,omitempty"`
	Borderhittest     bool     `xml:"borderhittest,attr,omitempty"`
	Subdiv            bool     `xml:"subdiv,attr,omitempty"`
	Ishotspot         bool     `xml:"ishotspot,attr,omitempty"`
	Loading           bool     `xml:"loading,attr,omitempty"`
	Loaded            bool     `xml:"loaded,attr,omitempty"`
	Loadedurl         string   `xml:"loadedurl,attr,omitempty"`
	Hovering          bool     `xml:"hovering,attr,omitempty"`
	Pressed           bool     `xml:"pressed,attr,omitempty"`
	Hitx              float64  `xml:"hitx,attr,omitempty"`
	Hity              float64  `xml:"hity,attr,omitempty"`
	Hitd              float64  `xml:"hitd,attr,omitempty"`
	Onover            string   `xml:"onover,attr,omitempty"`
	Onhover           string   `xml:"onhover,attr,omitempty"`
	Onout             string   `xml:"onout,attr,omitempty"`
	Ondown            string   `xml:"ondown,attr,omitempty"`
	Onup              string   `xml:"onup,attr,omitempty"`
	Onclick           string   `xml:"onclick,attr,omitempty"`
	Onloaded          string   `xml:"onloaded,attr,omitempty"`
}

type Events struct {
	XMLName              xml.Name `xml:"events"`
	Name                 string   `xml:"name,attr,omitempty"`
	Keep                 bool     `xml:"keep,attr,omitempty"`
	Onenterfullscreen    string   `xml:"onenterfullscreen,attr,omitempty"`
	Onexitfullscreen     string   `xml:"onexitfullscreen,attr,omitempty"`
	Onxmlcomplete        string   `xml:"onxmlcomplete,attr,omitempty"`
	Onpreviewcomplete    string   `xml:"onpreviewcomplete,attr,omitempty"`
	Onloadcomplete       string   `xml:"onloadcomplete,attr,omitempty"`
	Onblendcomplete      string   `xml:"onblendcomplete,attr,omitempty"`
	Onnewpano            string   `xml:"onnewpano,attr,omitempty"`
	Onremovepano         string   `xml:"onremovepano,attr,omitempty"`
	Onnewscene           string   `xml:"onnewscene,attr,omitempty"`
	Onxmlerror           string   `xml:"onxmlerror,attr,omitempty"`
	Onloaderror          string   `xml:"onloaderror,attr,omitempty"`
	Onkeydown            string   `xml:"onkeydown,attr,omitempty"`
	Onkeyup              string   `xml:"onkeyup,attr,omitempty"`
	Onclick              string   `xml:"onclick,attr,omitempty"`
	Onsingleclick        string   `xml:"onsingleclick,attr,omitempty"`
	Ondoubleclick        string   `xml:"ondoubleclick,attr,omitempty"`
	Onmousedown          string   `xml:"onmousedown,attr,omitempty"`
	Onmouseup            string   `xml:"onmouseup,attr,omitempty"`
	Onmousewheel         string   `xml:"onmousewheel,attr,omitempty"`
	Oncontextmenu        string   `xml:"oncontextmenu,attr,omitempty"`
	Onidle               string   `xml:"onidle,attr,omitempty"`
	Onviewchange         string   `xml:"onviewchange,attr,omitempty"`
	Onviewchanged        string   `xml:"onviewchanged,attr,omitempty"`
	Onresize             string   `xml:"onresize,attr,omitempty"`
	Onframebufferresize  string   `xml:"onframebufferresize,attr,omitempty"`
	Onautorotatestart    string   `xml:"onautorotatestart,attr,omitempty"`
	Onautorotatestop     string   `xml:"onautorotatestop,attr,omitempty"`
	Onautorotateoneround string   `xml:"onautorotateoneround,attr,omitempty"`
	Onautorotatechange   string   `xml:"onautorotatechange,attr,omitempty"`
	Oniphonefullscreen   string   `xml:"oniphonefullscreen,attr,omitempty"`
}

type Action struct {
	XMLName xml.Name `xml:"action"`
	Name    string   `xml:"name,attr,omitempty"`
	Type    string   `xml:"type,attr,omitempty"`
	Scope   string   `xml:"scope,attr,omitempty"`
	Args    string   `xml:"args,attr,omitempty"`
	Autorun string   `xml:"autorun,attr,omitempty"`
	Protect bool     `xml:"protect,attr,omitempty"`
	Secure  bool     `xml:"secure,attr,omitempty"`
	Content string   `xml:"content,attr,omitempty"`
}

type Contextmenu struct {
	XMLName     xml.Name `xml:"contextmenu"`
	Fullscreen  bool     `xml:"fullscreen,attr,omitempty"`
	Versioninfo bool     `xml:"versioninfo,attr,omitempty"`
	Touch       bool     `xml:"touch,attr,omitempty"`
	Customstyle string   `xml:"customstyle,attr,omitempty"`
	Enterfs     string   `xml:"enterfs,attr,omitempty"`
	Exitfs      string   `xml:"exitfs,attr,omitempty"`
}

type Network struct {
	XMLName    xml.Name `xml:"network"`
	Retrycount int      `xml:"retrycount,attr,omitempty"`
}

type Memory struct {
	XMLName xml.Name `xml:"memory"`
	Maxmem  int      `xml:"maxmem,attr,omitempty"`
}

type Security struct {
	XMLName xml.Name `xml:"security"`
	Cors    string   `xml:"cors,attr,omitempty"`
}

type Textstyle struct {
	XMLName xml.Name `xml:"textstyle"`
}

type Data struct {
	XMLName xml.Name `xml:"data"`
	Name    int      `xml:"name,attr,omitempty"`
	Content string   `xml:"content,attr,omitempty"`
}

type Scene struct {
	XMLName  xml.Name `xml:"scene"`
	Name     string   `xml:"name,attr,omitempty"`
	Title    string   `xml:"title,attr,omitempty"`
	Thumburl string   `xml:"thumburl,attr,omitempty"`
	Index    int      `xml:"index,attr,omitempty"`
	Onstart  string   `xml:"onstart,attr,omitempty"`
	Autoload bool     `xml:"autoload,attr,omitempty"`
	Content  string   `xml:"content,attr,omitempty"`

	// sub tags
	Include    []*Include  `xml:"include"`
	Preview    *Preview    `xml:"preview"`
	Image      *Image      `xml:"image"`
	View       *View       `xml:"view"`
	Area       *Area       `xml:"area"`
	Display    *Display    `xml:"display"`
	Control    *Control    `xml:"control"`
	Cursors    *Cursors    `xml:"cursors"`
	Autorotate *Autorotate `xml:"autorotate"`
	Plugin     []*Plugin   `xml:"plugin"`
	Hotspot    []*Hotspot  `xml:"hotspot"`
	Events     []*Events   `xml:"events"`
	Action     []*Action   `xml:"action"`
	Data       []*Data     `xml:"data"`
	Set        []*Set      `xml:"set"`
}

type Set struct {
	XMLName xml.Name `xml:"set"`
	Var     string   `xml:"var,attr,omitempty"`
	Val     string   `xml:"val,attr,omitempty"`
}
