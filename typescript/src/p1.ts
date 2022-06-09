function getInput(): string {
  return `forward 5
  down 5
  forward 8
  up 3
  down 8
  forward 2`
}

function parseLine(line: string): [number, number] {
  const [dir, a] = line.split(' ')
  const amount = +a // trick to convert to number
  let multiplier = 1
  let res: [number, number] = [0, 0]
  if (dir === 'up' || dir === 'backward') multiplier = -1
  if (dir === 'up' || dir === 'down') {
    res = [0, Number(amount) * multiplier]
  }
  if (dir === 'forward' || dir === 'backward') { 
    res = [Number(amount) * multiplier, 0]
  }

  return res
}

const items = getInput().split('\n').map(x => parseLine(x.trim()))
console.log(items)
const result = items.reduce((acc, ammount) => {
  acc[0] += ammount[0]
  acc[1] += ammount[1]
  return acc
}, [0, 0])

console.log(result, result[0] * result[1])
