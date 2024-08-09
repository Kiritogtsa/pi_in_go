package entries

type Salario struct {
	ID             int64
	Salariobruto   float64
	Ir             float64
	Inss           float64
	Adicional      float64
	Salarioliquido float64
	Mes            int
	Ano            int
	Decimo         float64
}

func Newsalario(salariovruto float64, adicional float64, mes int, ano int) *Salario {
	return &Salario{
		Salariobruto: salariovruto,
		Adicional:    adicional,
		Mes:          mes,
		Ano:          ano,
	}
}

func (s *Salario) calcInss() {
	switch {
	case s.Salariobruto <= 1302.00:
		s.Inss = s.Salariobruto * 0.075
	case s.Salariobruto <= 2571.29:
		s.Inss = 1302.00*0.075 + (s.Salariobruto-1302.00)*0.09
	case s.Salariobruto <= 3856.94:
		s.Inss = 1302.00*0.075 + (2571.29-1302.00)*0.09 + (s.Salariobruto-2571.29)*0.12
	case s.Salariobruto <= 7507.49:
		s.Inss = 1302.00*0.075 + (2571.29-1302.00)*0.09 + (3856.94-2571.29)*0.12 + (s.Salariobruto-3856.94)*0.14
	default:
		s.Inss = 1302.00*0.075 + (2571.29-1302.00)*0.09 + (3856.94-2571.29)*0.12 + (7507.49-3856.94)*0.14
	}
}

func (s *Salario) calcIr() {
	switch {
	case s.Salariobruto <= 1903.98:
		s.Ir = 0
	case s.Salariobruto <= 2826.65:
		s.Ir = (s.Salariobruto * 0.075) - 142.80
	case s.Salariobruto <= 3751.05:
		s.Ir = (s.Salariobruto * 0.15) - 354.80
	case s.Salariobruto <= 4664.68:
		s.Ir = (s.Salariobruto * 0.225) - 636.13
	default:
		s.Ir = (s.Salariobruto * 0.275) - 869.36
	}
}

func (s *Salario) CalcularLiquido() {
	s.calcInss()
	s.calcIr()
	s.Salarioliquido = s.Salariobruto - s.Inss - s.Ir + s.Adicional
}

func (s *Salario) CalcularDecimo() {
	s.Decimo = s.Salariobruto / 12
}
