<?php
namespace App;
/**
 * @param $input
 * @param $pad_length
 * @param $pad_string
 * @param $pad_type
 * 中英文混编时字符串表格输出
 * @return string
 */
function str_pad($input, $pad_length, $pad_string, $pad_type)
{
    $strlen = (strlen($input) + mb_strlen($input, 'UTF8')) / 2;
    if ($strlen < $pad_length) {
        $difference = $pad_length - $strlen;
        switch ($pad_type) {
            case STR_PAD_RIGHT:
                return $input . str_repeat($pad_string, $difference);
                break;
            case STR_PAD_LEFT:
                return str_repeat($pad_string, $difference) . $input;
                break;
            default:
                $left  = $difference / 2;
                $right = $difference - $left;
                return str_repeat($pad_string, $left) . $input . str_repeat($pad_string, $right);
                break;
        }
    } else {
        return $input;
    }
}

/**
 * @param $input
 * @param $pad_length
 * @param $pad_string
 * @param $pad_type
 * 中英文混编时字符串表格输出（适用于双字节填充符）
 *
 * @return string
 */
function str_pad_for_double_bytes($input, $pad_length, $pad_string, $pad_type = STR_PAD_BOTH)
{
    $strlen = (strlen($input) + mb_strlen($input, 'UTF8')) / 2;

    if ($strlen < $pad_length) {
        $difference = $pad_length - $strlen;
        switch ($pad_type) {
            case STR_PAD_RIGHT:
                $repeat_str = str_repeat($pad_string, floor($difference / 2));
                if ($difference % 2) {
                    $repeat_str = ' ' . $repeat_str;
                }
                return $input . $repeat_str;
                break;
            case STR_PAD_LEFT:
                $repeat_str = str_repeat($pad_string, floor($difference / 2));
                if ($difference % 2) {
                    $repeat_str .= ' ';
                }
                return $repeat_str . $input;
                break;
            default:
                //对于奇数字节数加个空格方便双字节处理
                $originalDiff = $difference;
                if ($difference % 2) {
                    $difference -= 1;//奇数减一使可被2整除
                }
                $left  = $difference / 2;
                $right = $difference - $left;
                if ($originalDiff % 2) {
                    $right += 1;
                }

                if ($left % 2 && $right % 2) {
                    $left  += 1;
                    $right -= 1;
                }
                //左侧字符串
                $left_repeat = str_repeat($pad_string, floor($left / 2));
                //右侧字符串
                $right_repeat = str_repeat($pad_string, floor($right / 2));

                if ($left % 2) {
                    $left_repeat .= ' ';
                }
                if ($right % 2) {
                    $right_repeat = ' ' . $right_repeat;
                }

                return $left_repeat . $input . $right_repeat;
                break;
        }
    } else {
        return $input;
    }
}
