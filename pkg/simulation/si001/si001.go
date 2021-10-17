package si001

import (
	"fmt"

	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/rsx/pkg/context"
	"github.com/xh3b4sd/rsx/pkg/step"
	"github.com/xh3b4sd/rsx/pkg/step/st002"
	"github.com/xh3b4sd/rsx/pkg/step/st003"
	"github.com/xh3b4sd/rsx/pkg/step/st004"
	"github.com/xh3b4sd/rsx/pkg/step/st005"
	"github.com/xh3b4sd/rsx/pkg/step/st006"
	"github.com/xh3b4sd/rsx/pkg/step/st007"
	"github.com/xh3b4sd/rsx/pkg/step/st008"
	"github.com/xh3b4sd/rsx/pkg/step/st009"
	"github.com/xh3b4sd/rsx/pkg/step/st010"
	"github.com/xh3b4sd/rsx/pkg/step/st011"
	"github.com/xh3b4sd/rsx/pkg/step/st012"
	"github.com/xh3b4sd/rsx/pkg/step/st013"
	"github.com/xh3b4sd/rsx/pkg/step/st014"
	"github.com/xh3b4sd/rsx/pkg/step/st015"
	"github.com/xh3b4sd/rsx/pkg/step/st016"
	"github.com/xh3b4sd/rsx/pkg/step/st017"
	"github.com/xh3b4sd/rsx/pkg/step/st018"
	"github.com/xh3b4sd/rsx/pkg/step/st019"
	"github.com/xh3b4sd/rsx/pkg/step/st020"
)

func Run() error {
	var err error

	ctx := context.Context{}

	steps := []step.Interface{

		//
		// initial network state
		//

		st002.Step{Index: 0 /*****/, Value: 1.00 /*********/, Comment: "mutate: set 1.00 DAI price floor"},
		st003.Step{Index: 1 /*****/, Value: 2.00 /*********/, Comment: "mutate: set 2.00 DAI price ceiling"},

		st004.Step{Index: 2 /*****/, Value: 3e06 /*********/, Comment: "mutate: add 3.0M DAI to treasury"},
		st017.Step{Index: 3 /*****/, Value: 3e06 /*********/, Comment: "verify: set 3.0M DAI in treasury"},
		st018.Step{Index: 4 /*****/, Value: 3e06 /*********/, Comment: "mutate: add 3.0M protocol debt in RSX"},
		st019.Step{Index: 5 /*****/, Value: 3e06 /*********/, Comment: "verify: set 3.0M protocol debt in RSX"},

		st005.Step{Index: 6 /*****/, Value: 4e06 /*********/, Comment: "mutate: add 4.0M RSX / DAI liquidity to pool"},
		st006.Step{Index: 7 /******************************/, Comment: "mutate: <amount> RSX circulating supply"},
		st008.Step{Index: 8 /*****/, Value: 1e06 /*********/, Comment: "verify: set 1.0M RSX circulating supply"},
		st007.Step{Index: 9 /******************************/, Comment: "mutate: <amount> RSX total supply"},
		st009.Step{Index: 10 /****/, Value: 4e06 /*********/, Comment: "verify: set 4.0M RSX total supply"},
		st011.Step{Index: 11 /*****************************/, Comment: "mutate: <amount> RSX market cap"},
		st010.Step{Index: 12 /****/, Value: 2e06 /*********/, Comment: "verify: set 2.0M RSX market cap"},

		st013.Step{Index: 13 /****/, Value: 2.00 /*********/, Comment: "verify: set RSX price of 2.00 DAI"},
		st020.Step{Index: 14 /*****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 15 /****/, Value: 0 /************/, Comment: "verify: set 0.0M excess reserves in treasury"},
		st016.Step{Index: 16 /****/, Value: 0 /************/, Comment: "verify: set 0.0M volume inflow in treasury"},

		//
		// start cycle 1
		//

		st012.Step{Index: 17 /****/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 18 /****/, Value: 2.21 /*********/, Comment: "verify: set RSX price of 2.21 DAI"},
		st015.Step{Index: 19 /****/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 20 /*****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 21 /****/, Value: 50e03 /********/, Comment: "verify: set 50k excess reserves in treasury"},
		st013.Step{Index: 22 /****/, Value: 1.99 /*********/, Comment: "verify: set RSX price of 1.99 DAI"},
		st016.Step{Index: 23 /****/, Value: 1e05 /*********/, Comment: "verify: set 100k volume inflow in treasury"},

		st012.Step{Index: 24 /****/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 25 /****/, Value: 2.00 /*********/, Comment: "verify: set RSX price of 2.00 DAI"},

		st012.Step{Index: 26 /****/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 27 /****/, Value: 2.20 /*********/, Comment: "verify: set RSX price of 2.20 DAI"},
		st015.Step{Index: 28 /****/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 29 /*****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 30 /****/, Value: 10e04 /********/, Comment: "verify: set 100k excess reserves in treasury"},
		st013.Step{Index: 31 /****/, Value: 1.99 /*********/, Comment: "verify: set RSX price of 1.99 DAI"},
		st016.Step{Index: 32 /****/, Value: 2e05 /*********/, Comment: "verify: set 200k volume inflow in treasury"},

		st012.Step{Index: 33 /****/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 34 /****/, Value: 2.00 /*********/, Comment: "verify: set RSX price of 2.00 DAI"},

		st012.Step{Index: 35 /****/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 36 /****/, Value: 2.20 /*********/, Comment: "verify: set RSX price of 2.20 DAI"},
		st015.Step{Index: 37 /****/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 38 /*****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 39 /****/, Value: 15e04 /********/, Comment: "verify: set 150k excess reserves in treasury"},
		st013.Step{Index: 40 /****/, Value: 1.99 /*********/, Comment: "verify: set RSX price of 1.99 DAI"},
		st016.Step{Index: 41 /****/, Value: 3e05 /*********/, Comment: "verify: set 300k volume inflow in treasury"},

		st012.Step{Index: 42 /****/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 43 /****/, Value: 2.00 /*********/, Comment: "verify: set RSX price of 2.00 DAI"},

		st012.Step{Index: 44 /****/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 45 /****/, Value: 2.20 /*********/, Comment: "verify: set RSX price of 2.20 DAI"},
		st015.Step{Index: 46 /****/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 47 /*****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 48 /****/, Value: 20e04 /********/, Comment: "verify: set 200k excess reserves in treasury"},
		st013.Step{Index: 49 /****/, Value: 1.99 /*********/, Comment: "verify: set RSX price of 1.99 DAI"},
		st016.Step{Index: 50 /****/, Value: 4e05 /*********/, Comment: "verify: set 400k volume inflow in treasury"},

		st012.Step{Index: 51 /****/, Value: 6000 /*********/, Comment: "mutate: buy RSX for 6k DAI from pool"},
		st013.Step{Index: 52 /****/, Value: 2.00 /*********/, Comment: "verify: set RSX price of 2.00 DAI"},

		st012.Step{Index: 53 /****/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 54 /****/, Value: 2.20 /*********/, Comment: "verify: set RSX price of 2.20 DAI"},
		st015.Step{Index: 55 /****/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 56 /*****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 57 /****/, Value: 25e04 /********/, Comment: "verify: set 250k excess reserves in treasury"},
		st013.Step{Index: 58 /****/, Value: 1.99 /*********/, Comment: "verify: set RSX price of 1.99 DAI"},
		st016.Step{Index: 59 /****/, Value: 5e05 /*********/, Comment: "verify: set 500k volume inflow in treasury"},

		st012.Step{Index: 60 /****/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 61 /****/, Value: 2.00 /*********/, Comment: "verify: set RSX price of 2.00 DAI"},

		st012.Step{Index: 62 /****/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 63 /****/, Value: 2.20 /*********/, Comment: "verify: set RSX price of 2.20 DAI"},
		st015.Step{Index: 64 /****/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 65 /*****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 66 /****/, Value: 30e04 /********/, Comment: "verify: set 300k excess reserves in treasury"},
		st013.Step{Index: 67 /****/, Value: 1.99 /*********/, Comment: "verify: set RSX price of 1.99 DAI"},
		st016.Step{Index: 68 /****/, Value: 6e05 /*********/, Comment: "verify: set 600k volume inflow in treasury"},

		st012.Step{Index: 69 /****/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 70 /****/, Value: 2.00 /*********/, Comment: "verify: set RSX price of 2.00 DAI"},

		st012.Step{Index: 71 /****/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 72 /****/, Value: 2.20 /*********/, Comment: "verify: set RSX price of 2.20 DAI"},
		st015.Step{Index: 73 /****/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 74 /*****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 75 /****/, Value: 35e04 /********/, Comment: "verify: set 350k excess reserves in treasury"},
		st013.Step{Index: 76 /****/, Value: 1.99 /*********/, Comment: "verify: set RSX price of 1.99 DAI"},
		st016.Step{Index: 77 /****/, Value: 7e05 /*********/, Comment: "verify: set 700k volume inflow in treasury"},

		st012.Step{Index: 78 /****/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 79 /****/, Value: 2.00 /*********/, Comment: "verify: set RSX price of 2.00 DAI"},

		st012.Step{Index: 80 /****/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 81 /****/, Value: 2.20 /*********/, Comment: "verify: set RSX price of 2.20 DAI"},
		st015.Step{Index: 82 /****/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 83 /*****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 84 /****/, Value: 40e04 /********/, Comment: "verify: set 400k excess reserves in treasury"},
		st013.Step{Index: 85 /****/, Value: 1.99 /*********/, Comment: "verify: set RSX price of 1.99 DAI"},
		st016.Step{Index: 86 /****/, Value: 8e05 /*********/, Comment: "verify: set 800k volume inflow in treasury"},

		st012.Step{Index: 87 /****/, Value: 6000 /*********/, Comment: "mutate: buy RSX for 6k DAI from pool"},
		st013.Step{Index: 88 /****/, Value: 2.00 /*********/, Comment: "verify: set RSX price of 2.00 DAI"},

		st012.Step{Index: 89 /****/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 90 /****/, Value: 2.20 /*********/, Comment: "verify: set RSX price of 2.20 DAI"},
		st015.Step{Index: 91 /****/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 92 /*****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 93 /****/, Value: 45e04 /********/, Comment: "verify: set 450k excess reserves in treasury"},
		st013.Step{Index: 94 /****/, Value: 1.99 /*********/, Comment: "verify: set RSX price of 1.99 DAI"},
		st016.Step{Index: 95 /****/, Value: 9e05 /*********/, Comment: "verify: set 900k volume inflow in treasury"},

		st012.Step{Index: 96 /****/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 97 /****/, Value: 2.00 /*********/, Comment: "verify: set RSX price of 2.00 DAI"},

		st012.Step{Index: 98 /****/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 99 /****/, Value: 2.20 /*********/, Comment: "verify: set RSX price of 2.20 DAI"},
		st015.Step{Index: 100 /***/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 101 /****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 102 /***/, Value: 50e04 /********/, Comment: "verify: set 500k excess reserves in treasury"},
		st013.Step{Index: 103 /***/, Value: 1.99 /*********/, Comment: "verify: set RSX price of 1.99 DAI"},
		st016.Step{Index: 104 /***/, Value: 1e06 /*********/, Comment: "verify: set 1.0M volume inflow in treasury"},

		st012.Step{Index: 105 /***/, Value: 6000 /*********/, Comment: "mutate: buy RSX for 6k DAI from pool"},
		st013.Step{Index: 106 /***/, Value: 2.00 /*********/, Comment: "verify: set RSX price of 2.00 DAI"},

		//
		// end cycle 1
		//

		st017.Step{Index: 107 /***/, Value: 2e06 /*********/, Comment: "verify: set 2.0M DAI in treasury"},
		st019.Step{Index: 108 /***/, Value: 3e06 /*********/, Comment: "verify: set 3.0M protocol debt in RSX"},
		st006.Step{Index: 109 /****************************/, Comment: "mutate: <amount> RSX circulating supply"},
		st008.Step{Index: 110 /***/, Value: 15e05 /********/, Comment: "verify: set 1.5M RSX circulating supply"},
		st007.Step{Index: 111 /****************************/, Comment: "mutate: <amount> RSX total supply"},
		st009.Step{Index: 112 /***/, Value: 45e05 /********/, Comment: "verify: set 4.5M RSX total supply"},
		st011.Step{Index: 113 /****************************/, Comment: "mutate: <amount> RSX market cap"},
		st010.Step{Index: 114 /***/, Value: 3e06 /*********/, Comment: "verify: set 3.0M RSX market cap"},

		st002.Step{Index: 115 /***/, Value: 1.30 /*********/, Comment: "mutate: set 1.30 DAI price floor"},
		st003.Step{Index: 116 /***/, Value: 2.60 /*********/, Comment: "mutate: set 2.60 DAI price ceiling"},
		st020.Step{Index: 117 /****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 118 /***/, Value: 50e03 /********/, Comment: "verify: set 50k excess reserves in treasury"},

		st017.Step{Index: 119 /***/, Value: 2e06 /*********/, Comment: "verify: set 2.0M DAI in treasury"},
		st019.Step{Index: 120 /***/, Value: 3e06 /*********/, Comment: "verify: set 3.0M protocol debt in RSX"},
		st006.Step{Index: 121 /****************************/, Comment: "mutate: <amount> RSX circulating supply"},
		st008.Step{Index: 122 /***/, Value: 15e05 /********/, Comment: "verify: set 1.5M RSX circulating supply"},
		st007.Step{Index: 123 /****************************/, Comment: "mutate: <amount> RSX total supply"},
		st009.Step{Index: 124 /***/, Value: 45e05 /********/, Comment: "verify: set 4.5M RSX total supply"},
		st011.Step{Index: 125 /****************************/, Comment: "mutate: <amount> RSX market cap"},
		st010.Step{Index: 126 /***/, Value: 3e06 /*********/, Comment: "verify: set 3.0M RSX market cap"},

		//
		// start cycle 2
		//

		st012.Step{Index: 127 /***/, Value: 380e03 /*******/, Comment: "mutate: buy RSX for 380k DAI from pool"},
		st013.Step{Index: 128 /***/, Value: 2.83 /*********/, Comment: "verify: set RSX price of 2.83 DAI"},
		st015.Step{Index: 129 /***/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 130 /****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 131 /***/, Value: 10e04 /********/, Comment: "verify: set 100k excess reserves in treasury"},
		st013.Step{Index: 132 /***/, Value: 2.59 /*********/, Comment: "verify: set RSX price of 2.59 DAI"},
		st016.Step{Index: 133 /***/, Value: 11e05 /********/, Comment: "verify: set 1.1M volume inflow in treasury"},

		st012.Step{Index: 134 /***/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 135 /***/, Value: 2.60 /*********/, Comment: "verify: set RSX price of 2.60 DAI"},

		st012.Step{Index: 136 /***/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 137 /***/, Value: 2.83 /*********/, Comment: "verify: set RSX price of 2.83 DAI"},
		st015.Step{Index: 138 /***/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 139 /****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 140 /***/, Value: 15e04 /********/, Comment: "verify: set 150k excess reserves in treasury"},
		st013.Step{Index: 141 /***/, Value: 2.59 /*********/, Comment: "verify: set RSX price of 2.59 DAI"},
		st016.Step{Index: 142 /***/, Value: 12e05 /********/, Comment: "verify: set 1.2M volume inflow in treasury"},

		st012.Step{Index: 143 /***/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 144 /***/, Value: 2.60 /*********/, Comment: "verify: set RSX price of 2.60 DAI"},

		st012.Step{Index: 145 /***/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 146 /***/, Value: 2.83 /*********/, Comment: "verify: set RSX price of 2.83 DAI"},
		st015.Step{Index: 147 /***/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 148 /****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 149 /***/, Value: 20e04 /********/, Comment: "verify: set 200k excess reserves in treasury"},
		st013.Step{Index: 150 /***/, Value: 2.59 /*********/, Comment: "verify: set RSX price of 2.59 DAI"},
		st016.Step{Index: 151 /***/, Value: 13e05 /********/, Comment: "verify: set 1.3M volume inflow in treasury"},

		st012.Step{Index: 152 /***/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 153 /***/, Value: 2.60 /*********/, Comment: "verify: set RSX price of 2.60 DAI"},

		st012.Step{Index: 154 /***/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 155 /***/, Value: 2.83 /*********/, Comment: "verify: set RSX price of 2.83 DAI"},
		st015.Step{Index: 156 /***/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 157 /****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 158 /***/, Value: 25e04 /********/, Comment: "verify: set 250k excess reserves in treasury"},
		st013.Step{Index: 159 /***/, Value: 2.59 /*********/, Comment: "verify: set RSX price of 2.59 DAI"},
		st016.Step{Index: 160 /***/, Value: 14e05 /********/, Comment: "verify: set 1.4M volume inflow in treasury"},

		st012.Step{Index: 161 /***/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 162 /***/, Value: 2.60 /*********/, Comment: "verify: set RSX price of 2.60 DAI"},

		st012.Step{Index: 163 /***/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 164 /***/, Value: 2.83 /*********/, Comment: "verify: set RSX price of 2.83 DAI"},
		st015.Step{Index: 165 /***/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 166 /****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 167 /***/, Value: 30e04 /********/, Comment: "verify: set 300k excess reserves in treasury"},
		st013.Step{Index: 168 /***/, Value: 2.59 /*********/, Comment: "verify: set RSX price of 2.59 DAI"},
		st016.Step{Index: 169 /***/, Value: 15e05 /********/, Comment: "verify: set 1.5M volume inflow in treasury"},

		st012.Step{Index: 170 /***/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 171 /***/, Value: 2.60 /*********/, Comment: "verify: set RSX price of 2.60 DAI"},

		st012.Step{Index: 172 /***/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 173 /***/, Value: 2.83 /*********/, Comment: "verify: set RSX price of 2.83 DAI"},
		st015.Step{Index: 174 /***/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 175 /****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 176 /***/, Value: 35e04 /********/, Comment: "verify: set 350k excess reserves in treasury"},
		st013.Step{Index: 177 /***/, Value: 2.59 /*********/, Comment: "verify: set RSX price of 2.59 DAI"},
		st016.Step{Index: 178 /***/, Value: 16e05 /********/, Comment: "verify: set 1.6M volume inflow in treasury"},

		st012.Step{Index: 179 /***/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 180 /***/, Value: 2.60 /*********/, Comment: "verify: set RSX price of 2.60 DAI"},

		st012.Step{Index: 181 /***/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 182 /***/, Value: 2.83 /*********/, Comment: "verify: set RSX price of 2.83 DAI"},
		st015.Step{Index: 183 /***/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 184 /****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 185 /***/, Value: 40e04 /********/, Comment: "verify: set 400k excess reserves in treasury"},
		st013.Step{Index: 186 /***/, Value: 2.59 /*********/, Comment: "verify: set RSX price of 2.59 DAI"},
		st016.Step{Index: 187 /***/, Value: 17e05 /********/, Comment: "verify: set 1.7M volume inflow in treasury"},

		st012.Step{Index: 188 /***/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 189 /***/, Value: 2.60 /*********/, Comment: "verify: set RSX price of 2.60 DAI"},

		st012.Step{Index: 190 /***/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 191 /***/, Value: 2.83 /*********/, Comment: "verify: set RSX price of 2.83 DAI"},
		st015.Step{Index: 192 /***/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 193 /****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 194 /***/, Value: 45e04 /********/, Comment: "verify: set 450k excess reserves in treasury"},
		st013.Step{Index: 195 /***/, Value: 2.59 /*********/, Comment: "verify: set RSX price of 2.59 DAI"},
		st016.Step{Index: 196 /***/, Value: 18e05 /********/, Comment: "verify: set 1.8M volume inflow in treasury"},

		st012.Step{Index: 197 /***/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 198 /***/, Value: 2.60 /*********/, Comment: "verify: set RSX price of 2.60 DAI"},

		st012.Step{Index: 199 /***/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 200 /***/, Value: 2.83 /*********/, Comment: "verify: set RSX price of 2.83 DAI"},
		st015.Step{Index: 201 /***/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 202 /****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 203 /***/, Value: 50e04 /********/, Comment: "verify: set 500k excess reserves in treasury"},
		st013.Step{Index: 204 /***/, Value: 2.59 /*********/, Comment: "verify: set RSX price of 2.59 DAI"},
		st016.Step{Index: 205 /***/, Value: 19e05 /********/, Comment: "verify: set 1.9M volume inflow in treasury"},

		st012.Step{Index: 206 /***/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 207 /***/, Value: 2.60 /*********/, Comment: "verify: set RSX price of 2.60 DAI"},

		st012.Step{Index: 208 /***/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 209 /***/, Value: 2.83 /*********/, Comment: "verify: set RSX price of 2.83 DAI"},
		st015.Step{Index: 210 /***/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 211 /****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 212 /***/, Value: 55e04 /********/, Comment: "verify: set 550k excess reserves in treasury"},
		st013.Step{Index: 213 /***/, Value: 2.59 /*********/, Comment: "verify: set RSX price of 2.59 DAI"},
		st016.Step{Index: 214 /***/, Value: 20e05 /********/, Comment: "verify: set 2.0M volume inflow in treasury"},

		st012.Step{Index: 215 /***/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 216 /***/, Value: 2.60 /*********/, Comment: "verify: set RSX price of 2.60 DAI"},

		st012.Step{Index: 217 /***/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 218 /***/, Value: 2.83 /*********/, Comment: "verify: set RSX price of 2.83 DAI"},
		st015.Step{Index: 219 /***/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 220 /****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 221 /***/, Value: 60e04 /********/, Comment: "verify: set 600k excess reserves in treasury"},
		st013.Step{Index: 222 /***/, Value: 2.59 /*********/, Comment: "verify: set RSX price of 2.59 DAI"},
		st016.Step{Index: 223 /***/, Value: 21e05 /********/, Comment: "verify: set 2.1M volume inflow in treasury"},

		st012.Step{Index: 224 /***/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 225 /***/, Value: 2.60 /*********/, Comment: "verify: set RSX price of 2.60 DAI"},

		st012.Step{Index: 226 /***/, Value: 1e05 /*********/, Comment: "mutate: buy RSX for 100k DAI from pool"},
		st013.Step{Index: 227 /***/, Value: 2.83 /*********/, Comment: "verify: set RSX price of 2.83 DAI"},
		st015.Step{Index: 228 /***/, Value: 1e05 /*********/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
		st020.Step{Index: 229 /****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 230 /***/, Value: 65e04 /********/, Comment: "verify: set 650k excess reserves in treasury"},
		st013.Step{Index: 231 /***/, Value: 2.59 /*********/, Comment: "verify: set RSX price of 2.59 DAI"},
		st016.Step{Index: 232 /***/, Value: 22e05 /********/, Comment: "verify: set 2.2M volume inflow in treasury"},

		st012.Step{Index: 233 /***/, Value: 4000 /*********/, Comment: "mutate: buy RSX for 4k DAI from pool"},
		st013.Step{Index: 234 /***/, Value: 2.60 /*********/, Comment: "verify: set RSX price of 2.60 DAI"},

		//
		// end cycle 2
		//

		st017.Step{Index: 235 /***/, Value: 32e05 /********/, Comment: "verify: set 3.2M DAI in treasury"},
		st019.Step{Index: 236 /***/, Value: 3e06 /*********/, Comment: "verify: set 3.0M protocol debt in RSX"},
		st006.Step{Index: 237 /****************************/, Comment: "mutate: <amount> RSX circulating supply"},
		st008.Step{Index: 238 /***/, Value: 196e04 /*******/, Comment: "verify: set 1.9M RSX circulating supply"},
		st007.Step{Index: 239 /****************************/, Comment: "mutate: <amount> RSX total supply"},
		st009.Step{Index: 240 /***/, Value: 496e04 /*******/, Comment: "verify: set 4.9M RSX total supply"},
		st011.Step{Index: 241 /****************************/, Comment: "mutate: <amount> RSX market cap"},
		st010.Step{Index: 242 /***/, Value: 509e04 /*******/, Comment: "verify: set 5.0M RSX market cap"},

		st002.Step{Index: 243 /***/, Value: 1.60 /*********/, Comment: "mutate: set 1.60 DAI price floor"},
		st003.Step{Index: 244 /***/, Value: 3.20 /*********/, Comment: "mutate: set 3.20 DAI price ceiling"},
		st020.Step{Index: 245 /****************************/, Comment: "mutate: <amount> excess reserves in treasury"},
		st014.Step{Index: 246 /***/, Value: 60e03 /********/, Comment: "verify: set 60k excess reserves in treasury"},

		st017.Step{Index: 247 /***/, Value: 32e05 /********/, Comment: "verify: set 3.2M DAI in treasury"},
		st019.Step{Index: 248 /***/, Value: 3e06 /*********/, Comment: "verify: set 3.0M protocol debt in RSX"},
		st006.Step{Index: 249 /****************************/, Comment: "mutate: <amount> RSX circulating supply"},
		st008.Step{Index: 250 /***/, Value: 196e04 /*******/, Comment: "verify: set 1.9M RSX circulating supply"},
		st007.Step{Index: 251 /****************************/, Comment: "mutate: <amount> RSX total supply"},
		st009.Step{Index: 252 /***/, Value: 496e04 /*******/, Comment: "verify: set 4.9M RSX total supply"},
		st011.Step{Index: 253 /****************************/, Comment: "mutate: <amount> RSX market cap"},
		st010.Step{Index: 254 /***/, Value: 509e04 /*******/, Comment: "verify: set 5.0M RSX market cap"},
	}

	for i, s := range steps {
		if s.Ind() != i {
			return tracer.Maskf(executionFailedError, "expected %d, got %d", i, s.Ind())
		}

		var spc string
		{
			if i >= 0 && i <= 9 {
				spc = "      "
			}
			if i >= 10 && i <= 99 {
				spc = "     "
			}
			if i >= 100 && i <= 999 {
				spc = "    "
			}
		}

		fmt.Printf("%d%s%s\n", s.Ind(), spc, s.Com())

		ctx, err = s.Run(ctx)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
