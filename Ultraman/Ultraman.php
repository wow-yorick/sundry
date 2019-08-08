<?php
/**
 * Ultraman 
 * 奥特曼
 * @package 
 * @version $id$
 * @copyright 2015-2016 The JIA Group
 * @author yorick <v@5zyx.com> 
 * @license Copyright © 2005-2015 www.jia.com All rights reserved
 */
class Ultraman {

    private static $count = 0;
    private static $attackCount = 0;
    private $hp = 0;
    private $mp = 0;
    private $name = '';
    private $hurt = 0;
    private $singleHurt = 0;
    private $singleStrongHurt = 0;
    private $multiHurt = 0;
    private $currentAttackType = '';

    public function __construct($name = 'Ultraman', $hp = 0, $mp = 0)
    {
        $this->hp = $hp;
        $this->mp = $mp;
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

    public function getMp()
    {
        return $this->mp;
    }

    public function deductMp($mp = 0)
    {
        if($this->mp < 20 && $mp < 0) {
            echo "mp is not enough!<br>".PHP_EOL;
            return false;
        }
        $this->mp = $this->mp + $mp;
        return true;
    }

    public function revertMp()
    {
        $this->mp = $this->mp + 5;
        echo "回复 5点MP!<br>".PHP_EOL;
    }

    public function getSingleHurt()
    {
        return $this->singleHurt;
    }

    private function setSingleHurt()
    {
        $this->singleHurt = rand(0, 30);
    }

    public function getSingleStrongHurt()
    {
        return $this->singleStrongHurt;
    }

    private function setSingleStrongHurt()
    {
        $this->singleStrongHurt = rand(0, 70);
    }

    public function getMultiHurt()
    {
        return $this->multiHurt;
    }

    private function setMultiHurt()
    {
        $this->multiHurt = rand(0, 20);
    }

    public function singleAttack(Monster $target)
    {
        //print_r($target);exit;
        $this->setSingleHurt();
        $hurt = $this->getSingleHurt();
        $target->deductHp($hurt);
        $this->hurt = $hurt;
        $this->currentAttackType = 'signle-attack';
        $this->calculateRet($target);
        return $this;
    }

    public function multiAttack(Array $targets)
    {
        foreach($targets as $target) {
            $this->setMultiHurt();
            $multiHurt = $this->getMultiHurt();
            $target->deductHp($multiHurt);
            $this->hurt = $multiHurt;
            $this->currentAttackType = 'multi-attack';
            $this->calculateRet($target);
        }

        return $this;

    }

    public function singleStrongAttack(Monster $target)
    {
        $setMpRet = $this->deductMp(-20);
        if(!$setMpRet) {
            return false;
        }
        $this->setSingleStrongHurt();
        $hurt = $this->getSingleStrongHurt();
        $target->deductHp($hurt);
        $this->hurt = $hurt;
        $this->currentAttackType = 'single-strong-attack';
        $this->calculateRet($target);
    }

    public function attack(Array $targets)
    {
        $proArr = array(
            'singleStrongAttack' => 10,
            'multiAttack' => 20,
            'singleAttack' => 70,
        );
        $attackType = $this->get_rand($proArr);

        if('singleStrongAttack' == $attackType && $ret = $this->singleStrongAttack(current($targets))) {
            $this->attack($targets);
            return $this;
        } elseif('multiAttack' == $attackType) {
            $this->multiAttack($targets);
        } elseif('singleAttack' == $attackType) {
            $this->singleAttack(current($targets));
        }
        return $this;
    }

    public function calculateRet(Monster $target)
    {
        $monName = $target->getName();
        $monHp = $target->getHp();

        $ultramanName = $this->getName();
        echo "Ultraman:{$ultramanName} attack[{$this->currentAttackType}] Monster:{$monName} hurt:{$this->hurt}<br>".PHP_EOL;
        echo "Monster:{$monName} status HP-{$monHp}<br>".PHP_EOL;
        if($monHp <= 0) {
            echo "Monster:{$monName} die<br>".PHP_EOL;
        }
    }

    private function get_rand($proArr) 
    {   
          
        $result = '';   
        //概率数组的总概率精度  
        $proSum = array_sum($proArr);   
        //概率数组循环    
        foreach ($proArr as $key => $proCur) {   
            $randNum = mt_rand(1, $proSum);               
            if ($randNum <= $proCur) {   
                $result = $key;                         
                break;   
            } else {   
                $proSum -= $proCur;                       
            }   
        }   
        unset ($proArr);   
        return $result;   
    }  
}
