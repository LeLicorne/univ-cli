package cmd

type Attack struct {
	Name   string
	Damage int
}

type Pokemon struct {
	Name    string
	HP      int
	MaxHP   int
	Attacks []Attack
}

func createPokemon() []Pokemon {
	return []Pokemon{
		{
			Name:  "Tortipouss",
			HP:    120,
			MaxHP: 120,
			Attacks: []Attack{
				{Name: "Fouet Lianes", Damage: 15},
				{Name: "Tranch'Herbe", Damage: 25},
				{Name: "Racines", Damage: 20},
				{Name: "Tempête Verte", Damage: 30},
			},
		},
		{
			Name:  "Ouisticram",
			HP:    110,
			MaxHP: 110,
			Attacks: []Attack{
				{Name: "Flammèche", Damage: 20},
				{Name: "Boutefeu", Damage: 30},
				{Name: "Griffe", Damage: 15},
				{Name: "Lance-Flammes", Damage: 25},
			},
		},
		{
			Name:  "Tiplouf",
			HP:    115,
			MaxHP: 115,
			Attacks: []Attack{
				{Name: "Pistolet à O", Damage: 20},
				{Name: "Bulles d'O", Damage: 25},
				{Name: "Écume", Damage: 15},
				{Name: "Hydrocanon", Damage: 30},
			},
		},
	}
}

func isFainted(p Pokemon) bool {
	return p.HP <= 0
}

func reduceHP(p *Pokemon, damage int) {
	p.HP -= damage
	if p.HP < 0 {
		p.HP = 0
	}
}
