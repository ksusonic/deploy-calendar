package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func Test_parseDoc(t *testing.T) {
	testDoc, err := html.Parse(strings.NewReader(testResponse()))
	if err != nil {
		t.Fatal(err)
	}

	actual, err := parseCalendarHTML(testDoc)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, actual, Calendar{
		ConditionsBySign: map[string]ConditionByDay{},
	})
}

func testResponse() string {
	return `
<!doctype html>
<html lang="ru">
<head>
    <title>Когда лучше деплоить в продакшен в ноябре 2024 года</title>
    <meta charset="utf-8">
    <link rel="stylesheet" href="/_css/style.css"/>
</head>
<body>
<div class="content">
    <div class="container">
        <div class="date">
            <div class="prev">
                <a href="/2024-10/">Октябрь</a>
            </div>
            <div class="next">
                <a href="/2024-12/">Декабрь</a>
            </div>
        </div>
        <h1>Когда лучше деплоить в продакшен в ноябре 2024 года</h1>
        <p>
            Информация адресована тем, кто планирует деплоить в прод в ближайшее время.
            Наиболее подходящие даты для этой процедуры всем знакам зодиака, смотрите таблицу.
        </p>
    </div>
    <div class="horizontal-scroll">
        <table>
            <thead>
            <tr>
                <td>Знак зодиака</td>
                <td>Благоприятные дни</td>
                <td>Нейтральные дни</td>
                <td>Неблагоприятные дни</td>
            </tr>
            </thead>
            <tbody>
            <tr class="row">
                <td><a href="/pdf/aries-2024-11/">Овен</a></td>
                <td class="calc">3 16 20 28</td>
                <td class="calc">1 6 7 10 12 14</td>
                <td class="calc">2 5 9 22 25 30</td>
            </tr>
            <tr class="row">
                <td><a href="/pdf/taurus-2024-11/">Телец</a></td>
                <td class="calc">3 7 13 16</td>
                <td class="calc">2 4 15 25</td>
                <td class="calc">9 20 22 23</td>
            </tr>
            <tr class="row">
                <td><a href="/pdf/gemini-2024-11/">Близнецы</a></td>
                <td class="calc">8 11 22</td>
                <td class="calc">6 7 17 21 25</td>
                <td class="calc">1 5 10 23 26 30</td>
            </tr>
            <tr class="row">
                <td><a href="/pdf/cancer-2024-11/">Рак</a></td>
                <td class="calc">5 15 20 23</td>
                <td class="calc">2 3 11 13 22 30</td>
                <td class="calc">7 9 12 17 18 19 26 28 29</td>
            </tr>
            <tr class="row">
                <td><a href="/pdf/leo-2024-11/">Лев</a></td>
                <td class="calc">20</td>
                <td class="calc">26 28</td>
                <td class="calc">5 10 12 19 25 30</td>
            </tr>
            <tr class="row">
                <td><a href="/pdf/virgo-2024-11/">Дева</a></td>
                <td class="calc">26 29</td>
                <td class="calc">2 11 14 20 24</td>
                <td class="calc">2 20</td>
            </tr>
            <tr class="row">
                <td><a href="/pdf/libra-2024-11/">Весы</a></td>
                <td class="calc">6 9</td>
                <td class="calc">3 5 7 11 17 18 20 25 28 30</td>
                <td class="calc">8 21 23</td>
            </tr>
            <tr class="row">
                <td><a href="/pdf/scorpio-2024-11/">Скорпион</a></td>
                <td class="calc">20 28 29</td>
                <td class="calc">10 16 18 30</td>
                <td class="calc">2 5 8 13 22 23 24 25 26 27</td>
            </tr>
            <tr class="row">
                <td><a href="/pdf/sagittarius-2024-11/">Стрелец</a></td>
                <td class="calc">2 3 6 10 19 20 21 26 30</td>
                <td class="calc">7 22 24 27 28</td>
                <td class="calc">6 13 21 26 29</td>
            </tr>
            <tr class="row">
                <td><a href="/pdf/capricorn-2024-11/">Козерог</a></td>
                <td class="calc">8 10 12</td>
                <td class="calc">15 22 30</td>
                <td class="calc">2 13 15 16 19 22 23 26 30</td>
            </tr>
            <tr class="row">
                <td><a href="/pdf/aquarius-2024-11/">Водолей</a></td>
                <td class="calc">20 22</td>
                <td class="calc">2 7 10 11 12 18 30</td>
                <td class="calc">1 14 15 16 17 19 23 26 29</td>
            </tr>
            <tr class="row">
                <td><a href="/pdf/pisces-2024-11/">Рыбы</a></td>
                <td class="calc">1 15 16 17 18 19 20 27 28 30</td>
                <td class="calc">4 6 7 8 9 12 14 21 23</td>
                <td class="calc">6 8 9 18 21</td>
            </tr>
            </tbody>
        </table>
    </div>
</div>
<footer><a href="/all/" class="no-print">Все месяцы</a>
    <br/>Проект сделал <a href="https://higimo.ru/">Хиги́мо</a>.
</footer>
</body>
</html>
`
}
