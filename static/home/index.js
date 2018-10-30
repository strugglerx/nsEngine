$("#wx").on("click", function(e){
	if($("#qr").css("display")=="none"){
		$("#qr").show();
	}else if($("#qr").css("display")=="block"){
		$("#qr").hide();
	}
	$("body").one("click", function(){
    	$("#qr").hide();
        
    });
    e.stopPropagation();
});
$("#qr").on("click", function(e){
    e.stopPropagation();
});

$("#page").on("click", function(e){
    if($("#qr").css("display")=="block"){    
    	$("#qr").hide();
    }
    e.stopPropagation();   
 });
    





