/**
 * Created by baidong on 2017/8/11.
 */
function emptyCheck() {
    var title = document.all['title'].value;
    if (title.length == 0) {
        alert("内容不能为空，请输入.")
        return false;
    }
    return true;
}
/*aliOss*/
$("#img_url").on("change", function(e) {
    var file = e.target.files[0]; //获取图片资源
    var filename = file.name;
    // 只选择图片文件
    if (!file.type.match('image.*')) {
        return false;
    }
    // LocalResizeIMG写法：
    lrz(file, {width: 200, fieldName: 'osstest'})
        .then(function (rst) {
            var ossData = new FormData();
            // 先请求授权，然后回调
            $.getJSON('/oss', function (json) {
                // 添加配置参数
                ossData.append('OSSAccessKeyId', json.accessid);
                ossData.append('policy', json.policy);
                ossData.append('Signature', json.signature);
                ossData.append('key', json.dir+filename);//目录+文件名
                ossData.append('success_action_status', 201); // 指定返回的状态码
                ossData.append('file', rst.file, filename);
                $.ajax({
                    url: json.host,
                    data: ossData,
                    dataType: 'xml', // 这里加个对返回内容的类型指定
                    processData: false,
                    contentType: false,
                    type: 'POST'
                }).done(function(data){
                    // 返回的上传信息
                    if ($(data).find('PostResponse')) {
                        var res = $(data).find('PostResponse');
                        console.info('Bucket：' + res.find('Bucket').text() );
                        console.info('Location：' + res.find('Location').text() );
                        console.info('Key：' + res.find('Key').text() );
                        console.info('ETag：' + res.find('ETag').text() );
                    }
                    // 图片预览
                    var img = new Image();
                    img.src = rst.base64;

                    img.onload = function () {
                        $(".preview_box").empty().append(img);
                    };
                });
            });
            return rst;
        })
        .catch(function (err) {
            // 万一出错了，这里可以捕捉到错误信息
            // 而且以上的then都不会执行
            alert('ERROR:'+err);
        })
        .always(function () {
            // 不管是成功失败，这里都会执行
        });

});