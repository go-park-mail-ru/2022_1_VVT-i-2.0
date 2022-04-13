package restaurants

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetRestaurants(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, `
[
  {
    "imgPath": "shokoladiza.png",
    "restName": "Шоколадница",
    "slug": "shokoladiza",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "smekalca_food",
    "restName": "Smekalca FooD",
    "slug": "smekalca_food",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "Subway",
    "restName": "subway",
    "slug": "subway.png",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "shaurma.png",
    "restName": "Шаурма",
    "slug": "shaurma",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "mac.png",
    "restName": "Макдональдс",
    "slug": "mac",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "KFC.png",
    "restName": "KFC",
    "slug": "KFC",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "BK.png",
    "restName": "Burger King",
    "slug": "BK",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "terem.png",
    "restName": "Теремок",
    "slug": "terem",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "zotman.png",
    "restName": "Zotmann Pizza",
    "slug": "zotman",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "tuktuk.png",
    "restName": "Tuk Tuk",
    "slug": "tuktuk",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "Bo.png",
    "restName": "BO",
    "slug": "Bo",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "paple.png",
    "restName": "Крошка картошка",
    "slug": "paple",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "",
    "restName": "",
    "slug": "",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "Якитория",
    "restName": "yaki.png",
    "slug": "yaki",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "dad.png",
    "restName": "Мама джанс",
    "slug": "dad",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "carlamov.png",
    "restName": "Варламов.сесть",
    "slug": "carlamov",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "allo.png",
    "restName": "Алло!Пицца",
    "slug": "allo",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
 {
    "imgPath": "fo89.png",
    "restName": "Fo 98",
    "slug": "fo89",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
},
{
  "imgPath": "pizzaexp.png",
  "restName": "Pizza Express 25/8",
  "slug": "pizzaexp",
  "timeToDeliver": "25-35 мин",
  "price": "500",
  "rating": "4.8"
},
  {
  "imgPath": "tanuki.png",
  "restName": "Tanuki",
  "slug": "tanuki",
  "timeToDeliver": "25-35 мин",
  "price": "500",
  "rating": "4.8"
},
  {
    "imgPath": "chay.png",
    "restName": "Чайона №2",
    "slug": "chay",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  },
  {
    "imgPath": "sakura.png",
    "restName": "Sakura",
    "slug": "sakura",
    "timeToDeliver": "25-35 мин",
    "price": "500",
    "rating": "4.8"
  }
]
`)
}

func GetDish(ctx echo.Context) error {
	slug := ctx.Param("slug")
	restaurants := map[string]string{
		"shokoladiza": "Шоколадница",
		"smekalca_food": "Smekalca FooD",
		"subway": "Subway",
		"shaurma": "Шаурма",
		"mac": "Макдональдс",
		"KFC": "KFC",
		"BK": "Burger King",
		"terem": "Теремок",
		"zotman": "Zotmann Pizza",
		"tuktuk": "Tuk Tuk",
		"Bo": "BO",
		"paple": "Крошка картошка",
		"yaki": "Якитория",
		"dad": "Мама джанс",
		"carlamov": "Варламов.сесть",
		"allo": "Алло!Пицца",
		"fo89": "Fo 98",
		"pizzaexp": "Pizza Express 25/8",
		"tanuki": "Tanuki",
		"chay": "Чайона №2",
		"sakura": "Sakura",
	}

	restname := restaurants[slug]

	return ctx.JSON(http.StatusOK, `{
  "restName": ` + restname +`,
  "products": [
    {
      "id": 1,
      "restaurant": 1,
      "name": "Баскет S с острыми крылышками",
      "description": "12 Крылышек в острой панировке.",
      "image_path": "basket_s.png",
      "calories": 100,
      "price": 500
    },
    {
      "id": 2,
      "restaurant": 1,
      "name": "Шефбургер острый",
      "description": "Попробуйте новый уникальный бургер от шефа! острая курочка в панировке Hot&spicy, сочные листья салата, пикантные маринованные огурчики, лук, фирменный соус Бургер и булочка с черно-белым кунжутом",
      "image_path": "shefburger.png",
      "calories": 100,
      "price": 200
    },
    {
      "id": 3,
      "restaurant": 1,
      "name": "Шефбургер Де Люкс острый",
      "description": "Острый бургер от шефа теперь де Люкс! Острое филе в хрустящей панировке, салат айсберг, маринованные огурцы, лук, фирменный соус Бургер, булочка с кунжутом, ломтик сыра и ломтик бекон.",
      "image_path": "shefburger_delux.png",
      "calories": 200,
      "price": 100
    },
    {
      "id": 4,
      "restaurant": 1,
      "name": "Шефбургер Де Люкс",
      "description": "Бургер от шефа теперь Де Люкс! Сочное филе в оригинальной панировке, томаты, салат айсберг, соус Цезарь, аппетитная булочка, ломтик сыра и ломтик бекона.",
      "image_path": "shefburger_delux_spaci.png",
      "calories": 300,
      "price": 900
    },
    {
      "id": 5,
      "restaurant": 1,
      "name": "Шефбургер Джуниор",
      "description": "Нежный соус Цезарь, два сочных стрипса в оригинальной панировке, салат айcберг и помидоры на пшеничной булочке с черно-белым кунжутом.",
      "image_path": "shefburger_junior.png",
      "calories": 400,
      "price": 800
    },
    {
      "id": 6,
      "restaurant": 1,
      "name": "Шефбургер Оригинальный",
      "description": "Попробуйте новый уникальный бургер от шефа! Нежный соус Цезарь, сочное филе в оригинальной панировке, салат айcберг и помидоры на пшеничной булочке с черно-белым кунжутом",
      "image_path": "shefburger_origin.png",
      "calories": 500,
      "price": 700
    },
    {
      "id": 7,
      "restaurant": 1,
      "name": "Картофель Фри Стандартный",
      "description": "Еще больше вкуса! В наших крупных ломтиках мы сохранили еще больше вкуса твоего любимого картофеля фри. Он получается именно таким, как ты любишь - с аппетитно хрустящей корочкой и мягкой, рассыпчатой серединкой. Любимое удовольствие!",
      "image_path": "patt_free.png",
      "calories": 600,
      "price": 600
    },
    {
      "id": 8,
      "restaurant": 1,
      "name": "Картофель по-деревенски",
      "description": "Рассыпчатый и хрустящий картофель по-деревенски - любимый вкус теперь в Kfc!",
      "image_path": "patt_country.png",
      "calories": 700,
      "price": 500
    },
    {
      "id": 9,
      "restaurant": 1,
      "name": "Сырные подушечки 5 штуки",
      "description": "Нежнейший сыр, обжаренный в хрустящей панировке! Идеальный перекус или дополнение к любимым блюдам!",
      "image_path": "chese_pat.png",
      "calories": 800,
      "price": 400
    },
    {
      "id": 10,
      "restaurant": 1,
      "name": "3 Острых стрипса",
      "description": "Только любимый вкус - и ничего лишнего. Потрясающе нежное куриное филе мы панируем вручную и готовим в ресторане по уникальному рецепту Полковника Сандерса. Совершенный вкус!",
      "image_path": "spaicy_strips.png",
      "calories": 900,
      "price": 300
    },
    {
      "id": 11,
      "restaurant": 1,
      "name": "3 Оригинальных стрипса",
      "description": "Только любимый вкус - и ничего лишнего. Потрясающе нежное куриное филе мы панируем вручную и готовим в ресторане по уникальному рецепту Полковника Сандерса. Совершенный вкус!",
      "image_path": "origin_strips.png",
      "calories": 100,
      "price": 200
    },
    {
      "id": 12,
      "restaurant": 1,
      "name": "Байтс малые",
      "description": "Любимые Байтс теперь с мягким вкусом! Свежие сочные кусочки курицы в золостистой хрустящей панировке.",
      "image_path": "bayts.png",
      "calories": 200,
      "price": 100
    },
    {
      "id": 13,
      "restaurant": 1,
      "name": "Печеная картошка с грибами и сыром",
      "description": "Картофель, сыр, зелень, грибной соус",
      "image_path": "1.jpeg",
      "calories": 100,
      "price": 500
    },
    {
      "id": 14,
      "restaurant": 1,
      "name": "Цезарь с куриной грудкой на гриле",
      "description": "Салат романо, салат айсберг, куриное филе, сыр пармезан, гренки, укроп, соус Цезарь",
      "image_path": "2.jpeg",
      "calories": 100,
      "price": 500
    },
    {
      "id": 15,
      "restaurant": 1,
      "name": "Котлета по-киевски",
      "description": "Куриная котлета, чесночное масло, зелень, картофельное пюре, помидоры черри, лук фри",
      "image_path": "3.jpeg",
      "calories": 100,
      "price": 500
    },
    {
      "id": 16,
      "restaurant": 1,
      "name": "Борщ",
      "description": "Говядина, картофель, морковь, свекла, капуста, сметана, лук репчатый",
      "image_path": "4.jpeg",
      "calories": 100,
      "price": 500
    },
    {
      "id": 17,
      "restaurant": 1,
      "name": "Куриная лапша",
      "description": "Курица, лапша яичная, морковь, укроп",
      "image_path": "5.jpeg",
      "calories": 100,
      "price": 500
    },
    {
      "id": 18,
      "restaurant": 1,
      "name": "Шашлык из свинины",
      "description": "Свинина, кинза, укроп, красный и репчатый лук",
      "image_path": "6.jpeg",
      "calories": 100,
      "price": 500
    },
    {
      "id": 19,
      "restaurant": 1,
      "name": "Картофель по-охотничьи",
      "description": "Золотистый жареный картофель с сочными свиными колбасками, томлеными вешенками и шампиньонами 350 гр",
      "image_path": "7.jpeg",
      "calories": 100,
      "price": 500
    },
    {
      "id": 20,
      "restaurant": 1,
      "name": "Калифорния с кунжутом",
      "description": "Краб-крем, огурцы, авокадо, шеф-соус, кунжут белый",
      "image_path": "8.jpeg",
      "calories": 100,
      "price": 500
    },
    {
      "id": 21,
      "restaurant": 1,
      "name": "Том Ям с морепродуктами",
      "description": "Креветки тигровые, мидии, бульон Том Ям, рис, шампиньоны, кокосовое молоко, сливки, кинза, томаты черри",
      "image_path": "9.jpeg",
      "calories": 100,
      "price": 500
    }
  ]
}`)
}
