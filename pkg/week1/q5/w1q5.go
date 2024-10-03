package w1q5

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

const (
	RestaurantInfoRegex = "^(.*) (\\d+)$"
	OrderInfoRegex      = "^(.*) (\\d+:\\d+) (\\d+:\\d+) (.*)$"
	OverlapQuery        = "overlap"
	ApprovalTime        = 5
	DeliveryTime        = 10
)

var scanner *bufio.Scanner

type Order struct {
	Customer         string
	Restaurant       *Restaurant
	CreatedAt        time.Time
	ReadyAt          time.Time
	NeedApprovalTime bool
}

type Restaurant struct {
	Name                 string
	EstimatedPreparation int
	Orders               []*Order
}

func read() string {
	scanner.Scan()
	return scanner.Text()
}

func getRegexMatches(str string, regularExpression string) []string {
	regexCompiler := regexp.MustCompile(regularExpression)
	matches := regexCompiler.FindStringSubmatch(str)
	return matches
}

func mustParseTime(q string) time.Time {
	t, err := time.Parse("15:04", q)
	if err != nil {
		panic(err)
	}
	return t
}

func scanRestaurants(count int, restaurants map[string]*Restaurant) error {
	for ; count > 0; count-- {
		restaurantMatches := getRegexMatches(read(), RestaurantInfoRegex)

		name := restaurantMatches[1]
		estimatedPreparation, err := strconv.Atoi(restaurantMatches[2])
		if err != nil {
			return err
		}

		restaurants[name] = &Restaurant{
			Name:                 name,
			EstimatedPreparation: estimatedPreparation,
			Orders:               make([]*Order, 0),
		}
	}
	return nil
}

func scanOrders(count int, restaurants map[string]*Restaurant) []*Order {
	allOrders := make([]*Order, 0)

	for ; count > 0; count-- {
		order, err := parseOrder(restaurants, read())
		if err != nil {
			return nil
		}

		if order.ReadyAt.Before(order.CreatedAt) {
			order.ReadyAt = order.ReadyAt.Add(time.Hour * 24)
		}

		allOrders = append(allOrders, order)
		restaurants[order.Restaurant.Name].Orders = append(restaurants[order.Restaurant.Name].Orders, order)
	}
	return allOrders
}

func parseOrder(restaurants map[string]*Restaurant, orderInfo string) (ord *Order, err error) {
	orderMatches := getRegexMatches(orderInfo, OrderInfoRegex)

	// unify to [Restaurant -> Customer] Form
	needApprovalTime := true
	idx1, idx2, idx3, idx4 := 1, 2, 3, 4
	if _, ok := restaurants[orderMatches[1]]; !ok {
		idx1, idx2, idx3, idx4 = idx4, idx3, idx2, idx1
		needApprovalTime = false
	}

	defer func() {
		if r := recover(); r != nil {
			ord, err = nil, r.(error)
		}
	}()

	return &Order{
		Restaurant:       restaurants[orderMatches[idx1]],
		Customer:         orderMatches[idx4],
		CreatedAt:        mustParseTime(orderMatches[idx2]),
		ReadyAt:          mustParseTime(orderMatches[idx3]),
		NeedApprovalTime: needApprovalTime,
	}, nil
}

func Solve(src io.Reader) (answer string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	scanner = bufio.NewScanner(src)
	restaurantsCount, err := strconv.Atoi(read())
	if err != nil {
		panic("")
	}

	allRestaurants := make(map[string]*Restaurant, restaurantsCount)
	err = scanRestaurants(restaurantsCount, allRestaurants)
	if err != nil {
		panic(err)
	}

	numOfOrders, err := strconv.Atoi(read())
	if err != nil {
		panic(err)
	}

	var allOrders []*Order
	if allOrders = scanOrders(numOfOrders, allRestaurants); allOrders == nil {
		panic("")
	}

	if query := read(); query == OverlapQuery {
		sort.Slice(allOrders, func(i, j int) bool {
			return allOrders[i].CreatedAt.Before(allOrders[j].CreatedAt)
		})

		maxOverlap := 0
		currentMax := 0
		var activeIntervals []time.Time

		for _, order := range allOrders {
			for len(activeIntervals) > 0 && !activeIntervals[0].After(order.CreatedAt) {
				activeIntervals = activeIntervals[1:]
				currentMax--
			}

			index := sort.Search(len(activeIntervals), func(i int) bool {
				return activeIntervals[i].After(order.ReadyAt)
			})
			activeIntervals = append(activeIntervals, time.Time{})
			copy(activeIntervals[index+1:], activeIntervals[index:])
			activeIntervals[index] = order.ReadyAt

			currentMax++

			if currentMax > maxOverlap {
				maxOverlap = currentMax
			}
		}
		answer = fmt.Sprintln(maxOverlap)

	} else {
		restaurant, exist := allRestaurants[query]
		if !exist {
			panic("")
		}

		counter := 0
		for _, order := range restaurant.Orders {
			totalDuration := order.ReadyAt.Sub(order.CreatedAt)
			totalMinutes := int(totalDuration.Minutes())

			var actualPreparation int
			if order.NeedApprovalTime {
				actualPreparation = totalMinutes - (ApprovalTime + DeliveryTime)
			} else {
				actualPreparation = totalMinutes - (DeliveryTime)
			}

			if actualPreparation >= restaurant.EstimatedPreparation {
				counter++
			}
		}
		answer = fmt.Sprintln(counter)
	}

	return
}

func init() {
	// make sure "os" package is present for submission script
	_ = os.DevNull
}
