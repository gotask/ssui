// mergely.go
package ssui

import (
	"strings"
)

var HtmlMergely = `<!DOCTYPE html>
<html>
<head>
  <meta http-equiv="content-type" content="text/html; charset=UTF-8">
  <title>Mergely demo</title>
  <meta http-equiv="content-type" content="text/html; charset=UTF-8">
  <meta name="robots" content="noindex, nofollow">
  <meta name="googlebot" content="noindex, nofollow">
  <meta name="viewport" content="width=device-width, initial-scale=1">

    <link rel="stylesheet" type="text/css" href="/css/result-light.css">

      <script type="text/javascript" src="//cdnjs.cloudflare.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
      <script type="text/javascript" src="//cdnjs.cloudflare.com/ajax/libs/codemirror/5.11.0/codemirror.min.js"></script>
      <script type="text/javascript" src="//cdn.rawgit.com/wickedest/Mergely/3.4.1/lib/mergely.js"></script>
      <link rel="stylesheet" type="text/css" href="//cdnjs.cloudflare.com/ajax/libs/codemirror/5.11.0/codemirror.min.css">
      <link rel="stylesheet" type="text/css" href="//cdn.rawgit.com/wickedest/Mergely/3.4.0/lib/mergely.css">

  <style id="compiled-css" type="text/css">
      h1, ul {
  margin: .3em 0;
}

.container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  margin: 0 .5em;
}

.diffs {
  flex: 1 1 auto;
  display: flex;
  flex-direction: column;
}
.diffs header * {
  display: inline-block;
  vertical-align: middle;
}
.diffs .compare-wrapper {
  flex: 1 1 auto;
  position: relative;
}
.diffs .compare-wrapper #compare {
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
}

/* Auto-height fix */
.mergely-column .CodeMirror {
  height: 100%;
}

  </style>

</head>
<body>
    <div class="container">

    <div class="diffs">
        <header>
            <h1>Visualized Diffs</h1>
            <button id="prev" title="Previous diff">▲</button>
            <button id="next" title="Next diff">▼</button>
            <button id="wrap" title="Toggle line wrapping">
                <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 10 10"><path d="M1 2h8M1 4.5h6A1 1 0 0 1 7 7H5v-.3l-.3.3.3.3V7M1 7h1.5" stroke-width="1.1" fill="none" stroke="#000"/></svg>
            </button>
        </header>
        <div class="compare-wrapper">
            <div id="compare">
            </div>
        </div>
    </div>
</div>

<script type="text/javascript">


var comp = $('#compare');

function downloadJSON(url, callback) {
  $.get(url, function(data) {
    var json = JSON.parse(data);
    callback(json.msg);
  });
}

comp.mergely({
  cmsettings: {
    readOnly: false,
    lineWrapping: true
  },
  wrap_lines: true,

  //Doesn't do anything?
  //autoresize: true,

  editor_width: 'calc(50% - 25px)',
  editor_height: '100%',

  lhs: function(setValue) {
    downloadJSON("/api/mergely?file=FILE_LEFT", setValue);
  },
  rhs: function(setValue) {
    downloadJSON("/api/mergely?file=FILE_RIGHT", setValue);
  }
});

function changeOptions(changer) {
  var options = comp.mergely('options');
  changer(options);

  comp.mergely('options', options);
  comp.mergely('update');
}

$('#prev').click(function() { comp.mergely('scrollToDiff', 'prev'); });
$('#next').click(function() { comp.mergely('scrollToDiff', 'next'); });
$('#wrap').click(function() { changeOptions(function(x) { x.wrap_lines = !x.wrap_lines; }); });

</script>
</body>
</html>
`

type OnGetFile func(user, file string) string

type HMergely struct {
	F OnGetFile
}

func (m *HMergely) Page(leftFileName, r string) string {
	s := strings.ReplaceAll(HtmlMergely, "FILE_LEFT", leftFileName)
	s = strings.ReplaceAll(s, "FILE_RIGHT", r)
	return s
}

var (
	mergely = &HMergely{}
)
