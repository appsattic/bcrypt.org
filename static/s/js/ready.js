$(function() {

  $('#generate-hash').click(function() {
    // get the params to send
    var params = {
      password : $('#password1').val(),
      cost     : $('#cost1').val(),
    };
    var paramStr = $.param(params, true);

    $.post('/api/generate-hash.json', paramStr, null, 'json')
      .done(function(data) {
        console.log('data:', data);
        $('#hash1').text(data.hash);
      })
      .fail(function(err) {
        console.warn('Something went wrong:', err);
      })
      .always(function() {
        console.log('Done');
      })
    ;
  });

  $('#check-password').click(function() {
    // get the params to send
    var params = {
      password : $('#password2').val(),
      hash     : $('#hash2').val(),
    };
    var paramStr = $.param(params, true);

    $.post('/api/check-password.json', paramStr, null, 'json')
      .done(function(data) {
        console.log('data:', data);
        if ( data.ok ) {
          $('#same').text('Yes');
        }
        else {
          $('#same').text('No');
        }
        $('#cost2').text(data.cost);
      })
      .fail(function(err) {
        console.warn('Something went wrong:', err);
      })
      .always(function() {
        console.log('Done');
      })
    ;
  });

});
