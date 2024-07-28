package main

type AdStore interface {
	GetAdsByCategory(categories []string) []Ad
}

type InMemoryAdStore struct {
	ads map[string][]Ad
}

func NewInMemorydStore() AdStore {
	ads := map[string][]Ad{
		"clothing": {
			{RedirectUrl: "/product/66VCHSJNUP", Text: "Tank top for sale. 20% off."},
		},
		"accessories": {
			{RedirectUrl: "/product/1YMWWN1N4O", Text: "Watch for sale. Buy one, get second kit for free"},
		},
		"footwear": {
			{RedirectUrl: "/product/L9ECAV7KIM", Text: "Loafers for sale. Buy one, get second one for free"},
		},
		"hair": {
			{RedirectUrl: "/product/2ZYFJ3GM2N", Text: "Hairdryer for sale. 50% off."},
		},
		"decor": {
			{RedirectUrl: "/product/0PUK6V6EV0", Text: "Candle holder for sale. 30% off."},
		},
		"kitchen": {
			{RedirectUrl: "/product/9SIQT8TOJO", Text: "Bamboo glass jar for sale. 10% off."},
			{RedirectUrl: "/product/6E92ZMYYFZ", Text: "Mug for sale. Buy two, get third one for free"},
		},
	}
	return InMemoryAdStore{
		ads: ads,
	}

}

func (store InMemoryAdStore) GetAdsByCategory(categories []string) []Ad {
	var ads []Ad
	for _, category := range categories {
		ads = append(ads, store.ads[category]...)
	}
	return ads
}
