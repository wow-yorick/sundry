// JavaScript Document
jquery radio取值，checkbox取值，select取值，radio选中，checkbox选中，select选中，及其相关 
获取一组radio被选中项的值
var item = $('input[@name=items][@checked]').val();
获取select被选中项的文本
var item = $("select[@name=items] option[@selected]").text();
select下拉框的第二个元素为当前选中值
$('#select_id')[0].selectedIndex = 1;
radio单选组的第二个元素为当前选中值
$('input[@name=items]').get(1).checked = true;

获取值：

文本框，文本区域：$("#txt").attr("value")；
多选框checkbox：$("#checkbox_id").attr("value")；
单选组radio：   $("input[@type=radio][@checked]").val();
下拉框select： $('#sel').val();

控制表单元素：
文本框，文本区域：$("#txt").attr("value",'');//清空内容
                 $("#txt").attr("value",'11');//填充内容

多选框checkbox： $("#chk1").attr("checked",'');//不打勾
                 $("#chk2").attr("checked",true);//打勾
                 if($("#chk1").attr('checked')==undefined) //判断是否已经打勾

单选组radio：    $("input[@type=radio]").attr("checked",'2');//设置value=2的项目为当前选中项
下拉框select：   $("#sel").attr("value",'-sel3');//设置value=-sel3的项目为当前选中项
                $("<option value='1'>1111</option><option value='2'>2222</option>").appendTo("#sel")//添加下拉框的option
                $("#sel").empty()；//清空下拉框

