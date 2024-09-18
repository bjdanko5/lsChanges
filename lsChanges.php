<?php
ini_set('display_errors', 1);
ini_set('default_socket_timeout', 15); // установить таймаут в 30 секунд
error_reporting(E_ALL);
session_start(); // Запускаем сессию
ini_set("soap.wsdl_cache_enabled", "0");
$wsdlLK = "http://192.168.10.128/zkh_lk1/ws/WebСервисLK?wsdl";
$options = [
  'login' => "Администратор",
  'password' => "",
  'trace' => 1,
  'exceptions' => 1,
  'connection_timeout' => 15
];
$logFile = 'lsChanges.log';

// Function to handle GET requests
function handleGetRequest()
{
  $id = isset($_GET['id']) ? $_GET['id'] : '';
  $base = isset($_GET['base']) ? $_GET['base'] : '';
  $dt = isset($_GET['dt']) ? $_GET['dt'] : '';
  $mode = isset($_GET['mode']) ? $_GET['mode'] : '';
  $start = isset($_GET['start']) ? $_GET['start'] : '';
  $end = isset($_GET['end']) ? $_GET['end'] : '';

 /*  if (!isset($id) || !isset($base) || !isset($dt) || !isset($mode)) {
    http_response_code(500);
    echo json_encode(['Ошибка' => 'Не указаны обязательные(id,base,dt,mode) параметры запроса.'], JSON_UNESCAPED_UNICODE);
    exit;
  }
 */
  // Validate input parameters
 /*  if (!isset($id) || !isset($base) || !isset($dt) || !isset($mode)) {
    http_response_code(400);
    echo json_encode(['Ошибка' => 'Не указаны обязательные(id,base,dt,mode) параметры запроса.'], JSON_UNESCAPED_UNICODE);
    exit;
  } */
  global $soapClient;
  global $responseData;
  //отладка
 /*  $id = '201000000038';
  $base = '04';
  $dt = '12.09.2024';
  $mode = 'status';
  $mode = 'changes';
  $start = '1';
  $end = '100'; */

  $result = $soapClient->lsChanges([
    'id' => $id,
    'base' => $base,
    'dt' => $dt,
    'mode' => $mode,
    'start' => $start,
    'end' => $end
  ])->return;
  $responce = json_decode($result, true);

  if (isset($responce['Ошибка']) && $responce['Ошибка']!='') {
    $responce['Ошибка'] = "<b>Error 500</b> ".$responce['Ошибка'];
    http_response_code(500);
    echo $responce['Ошибка'];
   // echo json_encode(['Ошибка' => 'Не указаны обязательные(id,base,dt,mode) параметры запроса.'], JSON_UNESCAPED_UNICODE);
    exit;
  }

  // Return response data in JSON format
  header('Content-Type: application/json; charset=UTF-8');
  //echo json_encode( $responseData,JSON_UNESCAPED_UNICODE );
  echo $result;
}
// Function to get the client ip address
function get_client_ip_server() {
  error_reporting(E_ALL & ~E_WARNING);
  $ipaddress = '';
  if (isset($_SERVER['HTTP_CLIENT_IP']) && $_SERVER['HTTP_CLIENT_IP'] !== '') {
      $ipaddress = $_SERVER['HTTP_CLIENT_IP'];
  } elseif (isset($_SERVER['HTTP_X_FORWARDED_FOR']) && $_SERVER['HTTP_X_FORWARDED_FOR'] !== '') {
      $ipaddress = $_SERVER['HTTP_X_FORWARDED_FOR'];
  } elseif (isset($_SERVER['HTTP_X_FORWARDED']) && $_SERVER['HTTP_X_FORWARDED'] !== '') {
      $ipaddress = $_SERVER['HTTP_X_FORWARDED'];
  } elseif (isset($_SERVER['HTTP_FORWARDED_FOR']) && $_SERVER['HTTP_FORWARDED_FOR'] !== '') {
      $ipaddress = $_SERVER['HTTP_FORWARDED_FOR'];
  } elseif (isset($_SERVER['HTTP_FORWARDED']) && $_SERVER['HTTP_FORWARDED'] !== '') {
      $ipaddress = $_SERVER['HTTP_FORWARDED'];
  } elseif (isset($_SERVER['REMOTE_ADDR']) && $_SERVER['REMOTE_ADDR'] !== '') {
      $ipaddress = $_SERVER['REMOTE_ADDR'];
  } else {
      $ipaddress = 'UNKNOWN';
  }
  error_reporting(E_ALL);
  return $ipaddress;
}
function requiredLogPart()
{
  // Получите текущую дату и время
  $dateTime = date('Y-m-d H:i:s');

  // Получите IP-адрес клиента
  $ipAddress = $_SERVER['REMOTE_ADDR'];
  $ipAddress = get_client_ip_server();

  // Получите GET-параметры
  $getParams = $_GET;

  // Создайте строку для логирования
  error_reporting(E_ALL & ~E_WARNING);
  $logString = "$dateTime | $ipAddress | " . json_encode($getParams). "| ";
  error_reporting(E_ALL);
  return $logString;
}
function appendToPrevLog($prevLogFile, $logLines)
{
  file_put_contents($prevLogFile, implode("\n", $logLines), FILE_APPEND);
}

function truncateLog($logFile, $maxLogSize, $numLinesToKeep)
{
  $logContent = file_get_contents($logFile);
  $logLines = explode("\n", $logContent);
  $logLines = array_slice($logLines, -$numLinesToKeep);
  $logContent = implode("\n", $logLines);
  file_put_contents($logFile, $logContent);
  return $logLines;
}

function logMessage($logFile, $msg = '')
{

  $maxLogSize = 8 * 1024;
  $maxPrevLogSize = 1 * 1024 * 1024;
  $prevLogFile = $logFile . '.prev';
  $numLinesToKeep = 100;


  $logString = requiredLogPart() . " " . $msg . "\n";
  $logSize = file_exists($logFile) ? filesize($logFile) : 0;
  if ($logSize > $maxLogSize) {
    $logLines = truncateLog($logFile, $maxLogSize, $numLinesToKeep);
    appendToPrevLog($prevLogFile, $logLines);
    truncateLog($prevLogFile, $maxPrevLogSize, 1000);
  }
  error_reporting(E_ALL & ~E_WARNING);
  try {
    $result = file_put_contents($logFile, $logString, FILE_APPEND | LOCK_EX);
    if ($result === false) {
      throw new Exception('Ошибка записи в лог файл lsChanges.log.Надо дать права на запись в каталог /lsChanges');
    }

  } catch (Exception $e) {
    $errorMessage = "Error: " . $e->getMessage();
    $error = array("Ошибка" => $errorMessage);
    header('Content-Type: application/json; charset=UTF-8');
    echo json_encode($error, JSON_UNESCAPED_UNICODE);
    //exit;
  }
  error_reporting(E_ALL);
}

// Handle GET request
logMessage($logFile, "STARTED");
if (isset($_GET['mode']) && $_GET['mode'] == 'log') {
    $log_file = $logFile;
    if (file_exists($log_file)) {
        $lines = file($log_file, FILE_IGNORE_NEW_LINES);
        $last_20_lines = array_slice($lines, -20);
        $last_20_lines = array_reverse($last_20_lines); // Reverse the order

 // Create a table header
 echo '<table border="1" cellpadding="5" cellspacing="0" style="border-collapse: collapse; width: 100%;">';
 echo '<caption style="font-weight: bold; font-size: 18px;">Последние записи Log</caption>';
 echo '<tr style="background-color: #f0f0f0;">';
 echo '<th style="padding: 10px;">Дата и время</th>';
 echo '<th style="padding: 10px;">IP-адрес</th>';
 echo '<th style="padding: 10px;">GET-параметры</th>';
 echo '<th style="padding: 10px;">Этап выполнения / Ошибка</th>';
 echo '</tr>';
 
 // Output each log entry as table cells
 foreach ($last_20_lines as $line) {
     $cells = explode('|', $line);
     echo '<tr style="background-color: #fff;">';
     foreach ($cells as $cell) {
         echo '<td style="padding: 10px; border: 1px solid #ddd;">' . $cell . '</td>';
     }
     echo '</tr>';
 }
 
 echo '</table>';
    } else {
        echo "Error: Log file does not exist.";
    }
    exit;
}
// SOAP client settings
try {
  $soapClient = new SoapClient($wsdlLK, $options);
  handleGetRequest();
} catch (Exception $e) {
  // Handle exceptions
  $errorMessage = $e instanceof SoapFault ? "SoapFault: " . $e->getMessage() : "Error: " . $e->getMessage();


  $error = array("Ошибка" => $errorMessage);
  header('Content-Type: application/json; charset=UTF-8');
  logMessage($logFile, $errorMessage);
  echo json_encode($error, JSON_UNESCAPED_UNICODE);

}

session_destroy();
logMessage($logFile, "COMPLETE");