package solution380

import "testing"

func TestCase(t *testing.T) {
	var set = Constructor()
	insert := set.Insert(1) // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
	if !insert {
		t.Errorf("insert 1 failed")
	}
	remove := set.Remove(2) // 返回 false ，表示集合中不存在 2 。
	if remove {
		t.Errorf("remove 2 failed")
	}

	b := set.Insert(2) // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
	if !b {
		t.Errorf("insert 2 failed")
	}
	r := set.GetRandom() // getRandom 应随机返回 1 或 2 。
	if r != 1 && r != 2 {
		t.Errorf("getRandom failed")
	}
	b2 := set.Remove(1) // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
	if !b2 {
		t.Errorf("remove 1 failed")
	}
	b3 := set.Insert(2) // 2 已在集合中，所以返回 false 。
	if b3 {
		t.Errorf("insert 2 failed")
	}
	r2 := set.GetRandom() // 由于集合现在只有 2 了，getRandom 应随机返回 2 。
	if r2 != 2 {
		t.Errorf("getRandom failed")
	}
}
