<?php
ini_set('display_errors', 1);
error_reporting(E_ALL);
session_start(); // Запускаем сессию
ini_set("soap.wsdl_cache_enabled", "0");
$wsdlLK = "http://192.168.10.128/zkh_lk1/ws/WebСервисLK?wsdl";
$options = [
  'login' => "Администратор",
  'password' => "",
  'trace' => 1,
  'exceptions' => 1
];
$responseData = [
  'КоличествоИзменений' => '0',
  'ДатаИзменений' => '12.09.2024'
];
// SOAP client settings
$soapClient = new SoapClient($wsdlLK, $options);

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
    echo json_encode(['Ошибка' => 'Некорректные параметры запроса.']);
    exit;
  }
  global $soapClient;
  global $responseData;
  //отладка
  $id ='201000000038';
  $base  ='4';
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

// Handle GET request
handleGetRequest();
