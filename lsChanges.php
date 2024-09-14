<?php
ini_set('display_errors', 1);
ini_set('default_socket_timeout', 5); // установить таймаут в 30 секунд
error_reporting(E_ALL);
session_start(); // Запускаем сессию
ini_set("soap.wsdl_cache_enabled", "0");
$wsdlLK = "http://192.168.10.128/zkh_lk1/ws/WebСервисLK?wsdl";
$options = [
  'login' => "Администратор",
  'password' => "",
  'trace' => 1,
  'exceptions' => 1,
  'connection_timeout' => 5
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
  // Validate input parameters
  if (!isset($id) || !isset($base) || !isset($dt) || !isset($mode)) {
    http_response_code(400);
    echo json_encode(['Ошибка' => 'Не указаны обязательные(id,base,dt,mode) параметры запроса.'],JSON_UNESCAPED_UNICODE);
    exit;
  }
  global $soapClient;
  global $responseData;
  //отладка
  $id ='201000000038';
  $base  ='04';
  $dt ='12.09.2024';
  $mode ='status';
  $mode ='changes';
  $start ='1';
  $end ='100';

  $result = $soapClient->lsChanges([
    'id' => $id,
    'base' => $base,
    'dt' => $dt,
    'mode' => $mode,
    'start' => $start,
    'end' => $end
  ])->return;
  
   // Return response data in JSON format
  header('Content-Type: application/json; charset=UTF-8');
  //echo json_encode( $responseData,JSON_UNESCAPED_UNICODE );
  echo $result;
}
 function requiredLogPart() {
  // Получите текущую дату и время
  $dateTime = date('Y-m-d H:i:s');

  // Получите IP-адрес клиента
  $ipAddress = $_SERVER['REMOTE_ADDR'];

  // Получите GET-параметры
  $getParams = $_GET;

  // Создайте строку для логирования
  $logString = "$dateTime | $ipAddress | " . json_encode($getParams) ;
  return $logString;
}
 function appendToPrevLog($prevLogFile, $logLines) {
  file_put_contents($prevLogFile, implode("\n", $logLines), FILE_APPEND);
}

 function truncateLog($logFile, $maxLogSize, $numLinesToKeep) {
  $logContent = file_get_contents($logFile);
  $logLines = explode("\n", $logContent);
  $logLines = array_slice($logLines, -$numLinesToKeep);
  $logContent = implode("\n", $logLines);
  file_put_contents($logFile, $logContent);
  return $logLines;
}

function logMessage($logFile, $msg ='') {
  
  $maxLogSize = 8 * 1024; 
  $maxPrevLogSize = 1 * 1024 * 1024; 
  $prevLogFile = $logFile . '.prev';
  $numLinesToKeep = 100;


  $logString = requiredLogPart() ." ". $msg . "\n";
  $logSize = file_exists($logFile) ? filesize($logFile) : 0;
  if ($logSize > $maxLogSize) {
      $logLines = truncateLog($logFile, $maxLogSize, $numLinesToKeep);
      appendToPrevLog($prevLogFile, $logLines);
      truncateLog($prevLogFile, $maxPrevLogSize, 1000);
  }

  file_put_contents($logFile, $logString, FILE_APPEND | LOCK_EX);
}

// Handle GET request
logMessage($logFile, "STARTED");
// SOAP client settings
try {
    $soapClient = new SoapClient($wsdlLK, $options);
    handleGetRequest();
} catch (Exception $e) {
    // Handle exceptions
    $errorMessage = $e instanceof SoapFault ? "SoapFault: " . $e->getMessage() : "Error: " . $e->getMessage();
    header('Content-Type: application/json; charset=UTF-8');
    $error = array("Ошибка" => $errorMessage);
    echo json_encode($error, JSON_UNESCAPED_UNICODE);
    logMessage($logFile, $errorMessage);
}

session_destroy(); 
logMessage($logFile, "COMPLETE");