package lib

type goodInfo struct {
	Id    string
	Name  string
	Img   string
	Price int // 分
	Num   int
}

func (g *goodInfo) Sum() int {
	return g.Price * g.Num
}

type Cart struct {
	goods []*goodInfo
}

func NewCart() *Cart {
	return &Cart{goods: make([]*goodInfo, 0)}
}

func (c *Cart) Items() []*goodInfo {
	return c.goods
}

func (c *Cart) remove(i int) *goodInfo {
	var info = c.goods[i]
	copy(c.goods[i:], c.goods[i+1:])
	c.goods = c.goods[:len(c.goods)-1]
	return info
}

func (c *Cart) unshift(g *goodInfo) {
	var tmp = make([]*goodInfo, len(c.goods)+1)
	tmp[0] = g
	copy(tmp[1:], c.goods)
	c.goods = tmp
}

func (c *Cart) Add(id, name, img string, price, num int) {
	var info *goodInfo
	for i, g := range c.goods {
		if g.Id == id {
			info = c.remove(i)
			break
		}
	}
	if info != nil {
		info.Num += num
	} else {
		info = &goodInfo{
			Id:    id,
			Name:  name,
			Img:   img,
			Price: price,
			Num:   num,
		}
	}
	c.unshift(info)
}

func (c *Cart) Del(id string) {
	for i, v := range c.goods {
		if v.Id == id {
			c.remove(i)
		}
	}
}

func (c *Cart) GetGoodsNum() int {
	var sum = 0
	for _, v := range c.goods {
		sum += v.Num
	}
	return sum
}

func (c *Cart) Clear() {
	c.goods = make([]*goodInfo, 0)
}

func (c *Cart) GetTotalMoney() int {
	var totalMoney = 0
	for _, v := range c.goods {
		totalMoney += v.Price * v.Num
	}
	return totalMoney
}

func (c *Cart) Incr(id string) {
	for _, v := range c.goods {
		if v.Id == id {
			v.Num++
			break
		}
	}
}

func (c *Cart) Decr(id string) {
	for _, v := range c.goods {
		if v.Id == id {
			if v.Num > 0 {
				v.Num--
			}
			break
		}
	}
}
