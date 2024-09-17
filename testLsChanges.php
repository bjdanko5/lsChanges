<?php
$header = "Тестовые запросы для сервиса lsChanges";
echo "<h2 style='text-align: center; margin-bottom: 20px;'>$header</h2>";

function constructUrl($mode, $params)
{
    $lsChanges_script_name = "lsChanges";
    $url = "http://" . $_SERVER['HTTP_HOST'] . $_SERVER['REQUEST_URI'];
    $full_url = str_replace(basename(__FILE__, ".php"), $lsChanges_script_name, $url) . "?" . $params;
    return $full_url;
}
//Экоград Азов
$params_status = "id=201000003125&base=04&dt=17.09.2024&mode=status";
$full_url_status = constructUrl("status", $params_status);

$params_changes = "id=201000003125&base=04&dt=17.09.2024&mode=changes&start=1&end=100";
$full_url_changes = constructUrl("changes", $params_changes);

$params_log = "mode=log";
$full_url_log = constructUrl("log", $params_log);

//Экоград Новочеркасск
$params_status1 = "id=201000003592&base=04&dt=17.09.2024&mode=status";
$full_url_status1 = constructUrl("status", $params_status1);

$params_changes1 = "id=201000003592&base=04&dt=17.09.2024&mode=changes&start=1&end=100";
$full_url_changes1 = constructUrl("changes", $params_changes1);

?>

<div
    style="width: 80%; margin: 40px auto; padding: 20px; background-color: #f9f9f9; border: 1px solid #ddd; border-radius: 10px; box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);">
    <h3 style="margin-top: 0;">Тестовые запросы</h3>
    <hr>
    <h4 style="margin-top: 0;">Экоград Азов</h4>
    <div style="display: flex; flex-wrap: wrap; justify-content: space-between;">
        <div style="width: 45%; margin: 20px;">
            <h5>Режим Статус (mode = status)</h5>
            <a href="<?php echo $full_url_status; ?>"
                style="display: block; padding: 10px; background-color: #337ab7; color: #fff; border: none; border-radius: 5px; cursor: pointer;"><?php echo $full_url_status; ?></a>
        </div>
        <div style="width: 45%; margin: 20px;">
            <h5>Режим Изменения (mode = changes)</h5>
            <a href="<?php echo $full_url_changes; ?>"
                style="display: block; padding: 10px; background-color: #337ab7; color: #fff; border: none; border-radius: 5px; cursor: pointer;"><?php echo $full_url_changes; ?></a>
        </div>
    </div>
    <h4 style="margin-top: 0;">Экоград Новочеркасск</h4>
    <div style="display: flex; flex-wrap: wrap; justify-content: space-between;">
        <div style="width: 45%; margin: 20px;">
            <h5>Режим Статус (mode = status)</h5>
            <a href="<?php echo $full_url_status1; ?>"
                style="display: block; padding: 10px; background-color: #337ab7; color: #fff; border: none; border-radius: 5px; cursor: pointer;"><?php echo $full_url_status1; ?></a>
        </div>
        <div style="width: 45%; margin: 20px;">
            <h5>Режим Изменения (mode = changes)</h5>
            <a href="<?php echo $full_url_changes1; ?>"
                style="display: block; padding: 10px; background-color: #337ab7; color: #fff; border: none; border-radius: 5px; cursor: pointer;"><?php echo $full_url_changes1; ?></a>
        </div>
    </div>
    <div style="width: 45%; margin: 20px;">
            <h5>Вывести лог (mode = log)</h5>
            <a id="log-link" href="<?php echo $full_url_log; ?>"
                style="display: block; padding: 10px; background-color: #337ab7; color: #fff; border: none; border-radius: 5px; cursor: pointer;">Вывести
                лог</a>
            <div id="log-content" style="margin-top: 20px;"></div>
        </div>

</div>
<script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
<script>
    $(document).ready(function () {
        $('#log-link').click(function (e) {
            e.preventDefault();
            $.ajax({
                type: 'GET',
                url: '<?php echo $full_url_log; ?>',
                success: function (data) {
                    $('#log-content').html(data);
                  /*   $('#log-content').css({
                        'border': '1px dashed #ccc',
                        'padding': '10px'
                    }); */
                }
            });
        });
    });
</script>