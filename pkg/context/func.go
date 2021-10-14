package context

func (p ContextPool) RSXPrice() float64 {
	var pri float64
	{
		var c float64

		if p.RSXDAI.RSX.Price != 0 {
			c++
			pri += p.RSXDAI.RSX.Price
		}

		if p.RSXOHM.RSX.Price != 0 {
			c++
			pri += p.RSXOHM.RSX.Price
		}

		pri /= c
	}

	return pri
}
