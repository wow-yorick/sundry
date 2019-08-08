<?php
header("Content-Type:text/html;charset=utf-8");
require 'Ultraman.php';
require 'Monster.php';

function main()
{
    $ultraman = new Ultraman('迪迦Ultraman', 500, 30);
    $monsters = array();
    for($i=1; $i <=20; $i++) {
        $mon = new Monster("monster-{$i}", 100);
        array_push($monsters, $mon);
    }
    $i = 1;
    while($ultraman->getHp() > 0 && count($monsters) > 0) {
        echo "当前轮数:".$i.'<br>'.PHP_EOL;
        //奥特曼先手 奥特曼比较厉害
        $ultraman->attack($monsters);
        //小怪兽攻击
        $mk = array_rand($monsters);
        $monsters[$mk]->attack($ultraman);
        if($ultraman->getHp() > 0) {
            $ultraman->revertMp();
        }

        $monsters = removeDieMonster($monsters);
        $i++;
    }
    echo "Ultraman: status:hp-{$ultraman->getHp()} mp-{$ultraman->getMp()}<br>".PHP_EOL;
    $monsCount = count($monsters);
    echo "Monster: status:count-{$monsCount}<br>".PHP_EOL;
}

function removeDieMonster($monsters)
{
    if(empty($monsters)) {
        echo "没有活着的小怪兽了!<br>".PHP_EOL;
        return array();
    }
    foreach($monsters as $key => $monster) {
        if($monster->getHp() <= 0) {
            unset($monsters[$key]);
            $dieName = $monster->getName();
            echo "{$dieName} is die<br>".PHP_EOL;
        }
    }
    //print_r($monsters);exit;
    sort($monsters);
    return $monsters;
}

main();
