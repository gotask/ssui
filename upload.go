package ssui

type HUpload struct {
	*ElemBase
	Text     string
	OnUpload UploadFile
}

type UploadFile func(name string, data []byte)

var HtmlUpload = `<div class="layui-upload">
<button type="button" class="layui-btn" id="{{.Id}}"><i class="layui-icon"></i>{{.Text}}</button>
<script>
	layui.use(['upload'], function () {
	     var $ = layui.jquery,
	         upload = layui.upload;

		upload.render({
			elem: '#{{.Id}}'
			,url: '/api/upload?event_id={{.Id}}&url_router={{.Rout}}'
			,accept: 'file' //普通文件
			//,exts: 'zip|rar|7z' //只允许上传压缩文件
			,choose: function(obj){
				obj.preview(function(index, file, result){
					$("#{{.Id}}").val(file.name);
				});
			}
			,progress: function(n, elem){
				var percent = n + '%'; //获取进度百分比
				$("#{{.Id}}").text("正在上传..."+percent);
			}
			,done: function(res){
				console.log(res);
				if(res.code==0){
					$("#{{.Id}}").text("上传成功");
				}else{
					$("#{{.Id}}").text("上传失败");
				}
			}
			,error: function(index, upload){
				$("#{{.Id}}").text("上传失败");
			}
		});
	});
</script>
</div>`

func NewUpload(id, text string, onupload UploadFile) *HUpload {
	u := &HUpload{newElem(id, "upload", HtmlUpload), text, onupload}
	u.self = u
	return u
}

func (u *HUpload) Clone() HtmlElem {
	nu := NewUpload(u.Id, u.Text, u.OnUpload)
	nu.ElemBase.clone(u.ElemBase)
	return nu
}
