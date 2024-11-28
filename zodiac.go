package main

// ZodiacSign источник: https://en.wikipedia.org/wiki/Western_astrology
type ZodiacSign int

const (
	// Aries Овен 21.03–20.04
	Aries ZodiacSign = iota
	// Taurus Телец 21.04–21.05
	Taurus
	// Gemini Близнецы 22.05–21.06
	Gemini
	// Cancer Рак 22.06–22.07
	Cancer
	// Leo Лев 23.07–21.08
	Leo
	// Virgo Дева 22.08–23.09
	Virgo
	// Libra Весы 24.09–23.10
	Libra
	// Scorpio Скорпион 24.10–22.11
	Scorpio
	// Sagittarius Стрелец 23.11–22.12
	Sagittarius
	// Capricorn Козерог 23.12–20.01
	Capricorn
	// Aquarius Водолей 21.01–19.02
	Aquarius
	// Pisces Рыбы 20.02–20.03
	Pisces
)
