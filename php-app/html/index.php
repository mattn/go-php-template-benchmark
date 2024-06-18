<html>
<body>
<ul>
<?php
require_once('/app/vendor/smarty/smarty/libs/Smarty.class.php');
$smarty = new Smarty();
$json = file_get_contents("input.json");
$data = json_decode($json, true);
$smarty->assign('data', $data, true);
$smarty->display("app.tpl");
?>
</ul>
</body>
</html>
