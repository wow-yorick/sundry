<?php
/**
 * Monster 
 * 怪兽
 * @package 
 * @version $id$
 * @copyright 2015-2016 The JIA Group
 * @author yorick <v@5zyx.com> 
 * @license Copyright © 2005-2015 www.jia.com All rights reserved
 */
class Monster {

    private static $count = 0;
    private $hp = 0;
    private $name = '';
    private $hurt = 0;

    public function __construct($name = 'tiny-monster', $hp = 0)
    {
        $this->hp = $hp;
        $this->name = $name;
        self::$count++;
    }

    public function getName()
    {
        return $this->name;
    }

    public function getHp()
    {
        return $this->hp;
    }

    public function deductHp($hurt)
    {
        $ret = $this->hp - $hurt;
        $this->hp = $ret < 0 ? 0 : $ret;
    }

    public function getHurt()
    {
        return $this->hurt;
    }

    private function setHurt()
    {
        $this->hurt = rand(0, 20);
    }

    /**
     * attack 
     * 攻击
     * @access public
     * @return void
     */
    public function attack(Ultraman $target)
    {
        $this->setHurt();
        $hurt = $this->getHurt();
        $target->deductHp($hurt);
        $this->calculateRet($target);
        return $this;
    }

    public function calculateRet(Ultraman $target)
    {
        $monsterName = $this->getName();

        $ultramanName = $target->getName();
        $ultramanHp = $target->getHp();
        $ultramanMp = $target->getMp();
        echo "Monster:{$monsterName} attack Ultraman:{$ultramanName} hurt:{$this->hurt}<br>".PHP_EOL;
        echo "Ultraman:{$ultramanName} status HP-{$ultramanHp} MP-{$ultramanMp}<br>".PHP_EOL;
        if($ultramanHp <= 0) {
            echo "ultraman:{$ultramanName} die<br>".PHP_EOL;
        }
    }
}
