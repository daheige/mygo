//对map的key/val交换位置
package hgMap

type Map_s_i map[string]int //自定义类型
type Map_i_s map[int]string

/**
 * 定义翻转map函数
 */
func Fclip(bar Map_s_i) Map_i_s {
	inv_map := make(Map_i_s, len(bar))
	for k, v := range bar {
		inv_map[v] = k
	}

	return inv_map
}
