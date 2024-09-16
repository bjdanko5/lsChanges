<?php
$header = "Тестовые запросы для сервиса lsChanges";
echo "<h3 style='text-align: center; margin-bottom: 20px;'>$header</h3>";

function constructUrl($mode, $params) {
    $lsChanges_script_name = "lsChanges";
    $url = "http://" . $_SERVER['HTTP_HOST'] . $_SERVER['REQUEST_URI'];
    $full_url = str_replace(basename(__FILE__, ".php"), $lsChanges_script_name, $url) . "?" . $params;
    return $full_url;
}

$params_status = "id=201000000038&base=04&dt=12.09.2024&mode=status";
$full_url_status = constructUrl("status", $params_status);

$params_changes = "id=201000000038&base=04&dt=12.09.2024&mode=changes&start=1&end=100";
$full_url_changes = constructUrl("changes", $params_changes);

?>

<div style="width: 80%; margin: 40px auto; padding: 20px; background-color: #f9f9f9; border: 1px solid #ddd; border-radius: 10px; box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);">
    <h4 style="margin-top: 0;">Тестовые запросы</h4>
    <hr>
    <div style="display: flex; flex-wrap: wrap; justify-content: space-between;">
        <div style="width: 45%; margin: 20px;">
            <h5>Режим Статус (mode = status)</h5>
            <a href="<?php echo $full_url_status; ?>" style="display: block; padding: 10px; background-color: #337ab7; color: #fff; border: none; border-radius: 5px; cursor: pointer;"><?php echo $full_url_status; ?></a>
        </div>
        <div style="width: 45%; margin: 20px;">
            <h5>Режим Изменения (mode = changes)</h5>
            <a href="<?php echo $full_url_changes; ?>" style="display: block; padding: 10px; background-color: #337ab7; color: #fff; border: none; border-radius: 5px; cursor: pointer;"><?php echo $full_url_changes; ?></a>
        </div>
    </div>
</div>