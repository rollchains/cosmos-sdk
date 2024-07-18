package types

// TODO: ideally get this into sdk types instead of the cosmossdk.io/store/types package

type multiplierGasMeter struct {
	basicGasMeter
	multiplierNumerator   uint64
	multiplierDenominator uint64
}

func NewMultiplierGasMeter(limit Gas, multiplierNumerator uint64, multiplierDenominator uint64) GasMeter {
	return &multiplierGasMeter{
		basicGasMeter: basicGasMeter{
			limit:    limit,
			consumed: 0,
			// lock:     &sync.Mutex{},
		},
		multiplierNumerator:   multiplierNumerator,
		multiplierDenominator: multiplierDenominator,
	}
}

func (g *multiplierGasMeter) adjustGas(original Gas) Gas {
	return original * g.multiplierNumerator / g.multiplierDenominator
}

func (g *multiplierGasMeter) ConsumeGas(amount Gas, descriptor string) {
	g.basicGasMeter.ConsumeGas(g.adjustGas(amount), descriptor)
}

func (g *multiplierGasMeter) RefundGas(amount Gas, descriptor string) {
	g.basicGasMeter.RefundGas(g.adjustGas(amount), descriptor)
}

func (g *multiplierGasMeter) Multiplier() (numerator uint64, denominator uint64) {
	return g.multiplierNumerator, g.multiplierDenominator
}

type infiniteMultiplierGasMeter struct {
	infiniteGasMeter
	multiplierNumerator   uint64
	multiplierDenominator uint64
}

func NewInfiniteMultiplierGasMeter(multiplierNumerator uint64, multiplierDenominator uint64) GasMeter {
	return &infiniteMultiplierGasMeter{
		infiniteGasMeter: infiniteGasMeter{
			consumed: 0,
			// lock:     &sync.Mutex{},
		},
		multiplierNumerator:   multiplierNumerator,
		multiplierDenominator: multiplierDenominator,
	}
}

func (g *infiniteMultiplierGasMeter) adjustGas(original Gas) Gas {
	return original * g.multiplierNumerator / g.multiplierDenominator
}

func (g *infiniteMultiplierGasMeter) ConsumeGas(amount Gas, descriptor string) {
	g.infiniteGasMeter.ConsumeGas(g.adjustGas(amount), descriptor)
}

func (g *infiniteMultiplierGasMeter) RefundGas(amount Gas, descriptor string) {
	g.infiniteGasMeter.RefundGas(g.adjustGas(amount), descriptor)
}

func (g *infiniteMultiplierGasMeter) Multiplier() (numerator uint64, denominator uint64) {
	return g.multiplierNumerator, g.multiplierDenominator
}

// Helpers for setting gas meter with parent ctx multiplier
func NewGasMeterWithMultiplier(ctxGasMeter GasMeter, limit uint64) GasMeter {
	if ctxGasMeter == nil {
		return NewGasMeter(limit)
	}
	n, d := ctxGasMeter.Multiplier()
	return NewMultiplierGasMeter(limit, n, d)
}

func NewInfiniteGasMeterWithMultiplier(ctxGasMeter GasMeter) GasMeter {
	if ctxGasMeter == nil {
		return NewInfiniteGasMeter()
	}
	n, d := ctxGasMeter.Multiplier()
	return NewInfiniteMultiplierGasMeter(n, d)
}
