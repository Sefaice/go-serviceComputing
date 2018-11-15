$(document).ready(function(){})

function submit(){
    var data = {};
    data['id'] =  $('.id-input').val()
    data['email'] =  $('.email-input').val()
    data['tel'] =  $('.tel-input').val()
    var postData = JSON.stringify(data)
    $.ajax({
        type: 'POST',
        url: '/submit',
        data: postData,
        contentType: 'application/json;charset=utf-8',
        dataType: 'json',
        timeout: 5000,
        success: function(result, xhr) {
            // success, show result 
            $('.form-result-hide').attr('class', 'form-result-show')
            $('.id-result').text($('.id-input').val())
            $('.email-result').text($('.email-input').val())
            $('.tel-result').text($('.tel-input').val())
            $('.serverId-result').text(result['serverId'])

            $('.id-input').val('')
            $('.email-input').val('')
            $('.tel-input').val('')
        },
        error: function(result, xhr) {
            //console.log(result)
            alert('服务器连接错误: ' + xhr)
        }
      })
}