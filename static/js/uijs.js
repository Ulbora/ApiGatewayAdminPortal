
// $(document)
// .ready(function() {

$(function () {
    $('.ui.menu .ui.dropdown').dropdown({
        on: 'hover'
    });
    $('.ui.menu a.item')
        .on('click', function () {
            $(this)
                .addClass('active')
                .siblings()
                .removeClass('active')
                ;
        })
        ;
});
// })
// ;

$(function () {
    $('a.item').click(function () {
        $('.item').removeClass('active');
        $(this).addClass('active');
    })
});


$(function () {
    $('.activating.element')
    .popup()
  ;
});



$(function () {
    $('.ui.radio.checkbox')
    .checkbox()
  ;
});

var urlId;
var urlName;
var showModel = function(id, name){  
    urlId = id;
    urlName = name;
    $(function(){
        $('.ui.modal').modal('show');
    });
    $(".ui.modal").modal({
		closable: true
    });
    document.getElementById("cont").innerHTML = "Delete URL " + name + "?";
}

var getUrlName = function(){    
    return urlName;
}
var getRedirectUriId = function(){
    return urlId;
}


// $(function () {
//     $('.ui.modal')
//     .modal()
// });

$(function(){
	$("#delete").click(function(){
		$(".ui.modal").modal('show');
	});
	$(".ui.modal").modal({
		closable: true
	});
});



var grantTypeId;
var grantType;
var showGrantTypeModel = function(id, gt){  
    grantTypeId = id;
    grantType = gt;
    $(function(){
        $('.ui.modal').modal('show');
    });
    $(".ui.modal").modal({
		closable: true
    });
    document.getElementById("cont").innerHTML = "Delete Grant Type " + gt + "?";
}

var getGrantTypeId = function(){
    return grantTypeId;
}



var roleId;
var clientRole;
var showRoleModel = function(id, rl){  
    roleId = id;
    clientRole = rl;
    $(function(){
        $('.ui.modal').modal('show');
    });
    $(".ui.modal").modal({
		closable: true
    });
    document.getElementById("cont").innerHTML = "Delete Client Role " + rl + "?";
}

var getRoleId = function(){
    return roleId;
}


var uriId;
var clientUir;
var allowedUriRoleId
var showAllowedUriModel = function(id, ri, rid){  
    uriId = id;
    clientUir = ri;
    allowedUriRoleId = rid;
    $(function(){
        $('.ui.modal').modal('show');
    });
    $(".ui.modal").modal({
		closable: true
    });
    document.getElementById("cont").innerHTML = "Delete Client URI " + ri + "?";
}

var getAllowedUriId = function(){
    return uriId;
}

var getAllowedUriRoleId = function(){
    return allowedUriRoleId;
}


var unHideAllowedUri = function(id){
    document.getElementById(id).disabled = false
    document.getElementById("del"+ id).style.display = 'none';
    document.getElementById("sub"+ id).style.display = 'block';
}