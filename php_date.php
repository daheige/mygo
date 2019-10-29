<?php
// php时间对象操作

$date = date_create("2019-10-15");
date_add($date, date_interval_create_from_date_string("100 days"));
echo date_format($date, "Y-m-d");
echo PHP_EOL;

$date = new DateTime("2019-10-15"); //创建一个date对象
//减去1天
$x = date_sub($date, date_interval_create_from_date_string("1 days"));

echo date_format($date, "Y-m-d");

/**
 * $ php t.php
2020-01-23
2019-10-14
 */
