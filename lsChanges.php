<?php

// SOAP client settings
$soapClient = new SoapClient('http://example.com/soap/wsdl');

// Function to handle GET requests
function handleGetRequest() {
  $id = $_GET['id'];
  $base = $_GET['base'];
  $dt = $_GET['dt'];
  $mode = $_GET['mode'];
  $start = $_GET['start'];
  $end = $_GET['end'];

  // Validate input parameters
  if (!isset($id) || !isset($base) || !isset($dt) || !isset($mode)) {
    http_response_code(400);
    echo json_encode(['error' => 'Invalid input parameters']);
    exit;
  }

  // Call SOAP method to retrieve data
  $result = $soapClient->lsChanges($id, $base, $dt, $mode, $start, $end);

  // Handle response data
  if ($mode == 'status') {
    $responseData = [
      'Статус' => [
        'КоличествоИзменений' => $result['КоличествоИзменений'],
        'ДатаИзменений' => $result['ДатаИзменений'],
      ],
    ];
  } elseif ($mode == 'changes') {
    $responseData = [
      'ДатаИзменений' => $result['ДатаИзменений'],
      'Изменения' => [],
    ];
    foreach ($result['Изменения'] as $change) {
      $responseData['Изменения'][] = [
        'ЛС' => $change['ЛС'],
        'Адрес' => $change['Адрес'],
        'ФИО' => $change['ФИО'],
        'Жильцов' => $change['Жильцов'],
        'Статус ЛС' => $change['СтатусЛС'],
        'Новый' => $change['Новый'],
      ];
    }
  }

  // Return response data in JSON format
  header('Content-Type: application/json; charset=UTF-8');
  echo json_encode($responseData);
}

// Handle GET request
handleGetRequest();

?>
/*
SOAP method lsChanges returns an associative array with the following structure:
[
  'КоличествоИзменений' => 100,
  'ДатаИзменений' => '22.08.2024',
  'Изменения' => [
    [
      'ЛС' => '1001',
      'Адрес' => 'г.Париж,ул.Сен-Дюбуа,д.5 кв.4',
      'ФИО' => 'Макрон Э.М.',
      'Жильцов' => 2,
      'СтатусЛС' => 'Открыт',
      'Новый' => 'Да',
    ],
    [
      'ЛС' => '1002',
      'Адрес' => 'г.Париж,ул.Сен-Дюбуа,д.5 кв.6',
      'ФИО' => 'Макрон Б.Б.',
      'Жильцов' => 3,
      'СтатусЛС' => 'Закрыт',
      'Новый' => 'Нет',
    ],
  ],
]
*/
/*
Описание сервиса

Сервис предназначен для предоставления информации о статусе и изменениях в личных счетах (ЛС) по запросу. Сервис взаимодействует с внешним SOAP-сервисом для получения необходимой информации.

Функциональность

Сервис предоставляет следующие возможности:

Получение информации о статусе ЛС по идентификатору ЛС, базе и дате.
Получение информации об изменениях в ЛС по идентификатору ЛС, базе, дате и режиму.
Режимы работы

Сервис поддерживает два режима работы:

Статус: предоставляет информацию о статусе ЛС, включая количество изменений и дату последнего изменения.
Изменения: предоставляет информацию об изменениях в ЛС, включая список изменений с указанием ЛС, адреса, ФИО, количества жильцов, статуса ЛС и признака нового изменения.
Входные данные

Сервис принимает следующие входные данные:

Идентификатор ЛС
База
Дата
Режим (статус или изменения)
Начало и конец диапазона изменений (для режима изменений)
Выходные данные

Сервис возвращает информацию в формате JSON, содержащую следующую информацию:

Для режима статуса: количество изменений, дата последнего изменения.
Для режима изменений: список изменений с указанием ЛС, адреса, ФИО, количества жильцов, статуса ЛС и признака нового изменения.
Требования к входным данным

Входные данные должны соответствовать следующим требованиям:

Идентификатор ЛС: обязательный, строковый.
База: обязательный, строковый.
Дата: обязательный, строковый в формате дд.мм.гггг.
Режим: обязательный, строковый (статус или изменения).
Начало и конец диапазона изменений: необязательные, целые числа (для режима изменений).
*/

//без ассоциативных массивов
/*

// SOAP client settings
$soapClient = new SoapClient('http://example.com/soap/wsdl');

// Function to handle GET requests
function handleGetRequest() {
  $id = $_GET['id'];
  $base = $_GET['base'];
  $dt = $_GET['dt'];
  $mode = $_GET['mode'];
  $start = $_GET['start'];
  $end = $_GET['end'];

  // Validate input parameters
  if (!isset($id) || !isset($base) || !isset($dt) || !isset($mode)) {
    http_response_code(400);
    echo json_encode(['error' => 'Invalid input parameters']);
    exit;
  }

  // Call SOAP method to retrieve data
  $result = $soapClient->lsChanges($id, $base, $dt, $mode, $start, $end);

  // Handle response data
  if ($mode == 'status') {
    $responseData = [
      'Статус' => [
        'КоличествоИзменений' => $result->КоличествоИзменений,
        'ДатаИзменений' => $result->ДатаИзменений,
      ],
    ];
  } elseif ($mode == 'changes') {
    $responseData = [
      'ДатаИзменений' => $result->ДатаИзменений,
      'Изменения' => [],
    ];
    foreach ($result->Изменения as $change) {
      $responseData['Изменения'][] = [
        'ЛС' => $change->ЛС,
        'Адрес' => $change->Адрес,
        'ФИО' => $change->ФИО,
        'Жильцов' => $change->Жильцов,
        'Статус ЛС' => $change->СтатусЛС,
        'Новый' => $change->Новый,
      ];
    }
  }

  // Return response data in JSON format
  header('Content-Type: application/json; charset=UTF-8');
  echo json_encode($responseData);
}

// Handle GET request
handleGetRequest();

?>
stdClass Object
(
  [КоличествоИзменений] => 100
  [ДатаИзменений] => 22.08.2024
  [Изменения] => Array
    (
      [0] => stdClass Object
        (
          [ЛС] => 1001
          [Адрес] => г.Париж,ул.Сен-Дюбуа,д.5 кв.4
          [ФИО] => Макрон Э.М.
          [Жильцов] => 2
          [СтатусЛС] => Открыт
          [Новый] => Да
        )
      [1] => stdClass Object
        (
          [ЛС] => 1002
          [Адрес] => г.Париж,ул.Сен-Дюбуа,д.5 кв.6
          [ФИО] => Макрон Б.Б.
          [Жильцов] => 3
          [СтатусЛС] => Закрыт
          [Новый] => Нет
        )
    )
)
*/ 