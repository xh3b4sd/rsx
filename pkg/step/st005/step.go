package st005

import (
	"math"

	"github.com/xh3b4sd/rsx/pkg/context"
)

type Step struct {
	Comment string
	Value   float64
}

func (s Step) Com() string {
	return s.Comment
}

// add <Value> RSX / DAI liquidity to pool
func (s Step) Run(ctx context.Context) (context.Context, error) {
	val := s.Value / 2

	ctx.Pool.RSXDAI.RSX.Amount = val / 2
	ctx.Pool.RSXDAI.RSX.Price = 2
	ctx.Pool.RSXDAI.RSX.Value = val

	ctx.Pool.RSXDAI.DAI.Amount = val / 1
	ctx.Pool.RSXDAI.DAI.Price = 1
	ctx.Pool.RSXDAI.DAI.Value = val

	ctx.Pool.RSXDAI.ConstantK = ctx.Pool.RSXDAI.RSX.Amount * ctx.Pool.RSXDAI.DAI.Amount
	ctx.Pool.RSXDAI.Liquidity = math.Sqrt(ctx.Pool.RSXDAI.RSX.Amount * ctx.Pool.RSXDAI.DAI.Amount)

	ctx.Protocol.Debt.RSX.Amount = s.Value / 2
	ctx.Protocol.Debt.RSX.Price = 2
	ctx.Protocol.Debt.RSX.Value = s.Value

	return ctx, nil
}

// // this low-level function should be called from a contract which performs important safety checks
// function mint(address to) external lock returns (uint liquidity) {
//     (uint112 _reserve0, uint112 _reserve1,) = getReserves(); // gas savings
//     uint balance0 = IERC20Uniswap(token0).balanceOf(address(this));
//     uint balance1 = IERC20Uniswap(token1).balanceOf(address(this));
//     uint amount0 = balance0.sub(_reserve0);
//     uint amount1 = balance1.sub(_reserve1);

//     bool feeOn = _mintFee(_reserve0, _reserve1);
//     uint _totalSupply = totalSupply; // gas savings, must be defined here since totalSupply can update in _mintFee
//     if (_totalSupply == 0) {
//         address migrator = IUniswapV2Factory(factory).migrator();
//         if (msg.sender == migrator) {
//             liquidity = IMigrator(migrator).desiredLiquidity();
//             require(liquidity > 0 && liquidity != uint256(-1), "Bad desired liquidity");
//         } else {
//             require(migrator == address(0), "Must not have migrator");
//             liquidity = Math.sqrt(amount0.mul(amount1)).sub(MINIMUM_LIQUIDITY);
//             _mint(address(0), MINIMUM_LIQUIDITY); // permanently lock the first MINIMUM_LIQUIDITY tokens
//         }
//     } else {
//         liquidity = Math.min(amount0.mul(_totalSupply) / _reserve0, amount1.mul(_totalSupply) / _reserve1);
//     }
//     require(liquidity > 0, 'UniswapV2: INSUFFICIENT_LIQUIDITY_MINTED');
//     _mint(to, liquidity);

//     _update(balance0, balance1, _reserve0, _reserve1);
//     if (feeOn) kLast = uint(reserve0).mul(reserve1); // reserve0 and reserve1 are up-to-date
//     emit Mint(msg.sender, amount0, amount1);
// }

// // this low-level function should be called from a contract which performs important safety checks
// function swap(uint amount0Out, uint amount1Out, address to, bytes calldata data) external lock {
//     require(amount0Out > 0 || amount1Out > 0, 'UniswapV2: INSUFFICIENT_OUTPUT_AMOUNT');
//     (uint112 _reserve0, uint112 _reserve1,) = getReserves(); // gas savings
//     require(amount0Out < _reserve0 && amount1Out < _reserve1, 'UniswapV2: INSUFFICIENT_LIQUIDITY');

//     uint balance0;
//     uint balance1;
//     { // scope for _token{0,1}, avoids stack too deep errors
//     address _token0 = token0;
//     address _token1 = token1;
//     require(to != _token0 && to != _token1, 'UniswapV2: INVALID_TO');
//     if (amount0Out > 0) _safeTransfer(_token0, to, amount0Out); // optimistically transfer tokens
//     if (amount1Out > 0) _safeTransfer(_token1, to, amount1Out); // optimistically transfer tokens
//     if (data.length > 0) IUniswapV2Callee(to).uniswapV2Call(msg.sender, amount0Out, amount1Out, data);
//     balance0 = IERC20Uniswap(_token0).balanceOf(address(this));
//     balance1 = IERC20Uniswap(_token1).balanceOf(address(this));
//     }
//     uint amount0In = balance0 > _reserve0 - amount0Out ? balance0 - (_reserve0 - amount0Out) : 0;
//     uint amount1In = balance1 > _reserve1 - amount1Out ? balance1 - (_reserve1 - amount1Out) : 0;
//     require(amount0In > 0 || amount1In > 0, 'UniswapV2: INSUFFICIENT_INPUT_AMOUNT');
//     { // scope for reserve{0,1}Adjusted, avoids stack too deep errors
//     uint balance0Adjusted = balance0.mul(1000).sub(amount0In.mul(3));
//     uint balance1Adjusted = balance1.mul(1000).sub(amount1In.mul(3));
//     require(balance0Adjusted.mul(balance1Adjusted) >= uint(_reserve0).mul(_reserve1).mul(1000**2), 'UniswapV2: K');
//     }

//     _update(balance0, balance1, _reserve0, _reserve1);
//     emit Swap(msg.sender, amount0In, amount1In, amount0Out, amount1Out, to);
// }
