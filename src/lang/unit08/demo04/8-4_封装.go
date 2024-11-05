package main

import "fmt"

type Computer struct {
	brand string
	price float32
}

func (c *Computer) SetBrand(brand string) {
	c.brand = brand
}

func (c *Computer) GetBrand() string {
	return c.brand
}

func (c *Computer) SetPrice(price float32) {
	c.price = price
}

func (c *Computer) GetPrice() float32 {
	return c.price
}

func (c Computer) String() string {
	return fmt.Sprintf("电脑：品牌=%s 价格=%v元", c.brand, c.price)
}

func main() {
	c := Computer{}
	c.SetBrand("小米")
	c.SetPrice(6000)
	fmt.Println(c)
}
