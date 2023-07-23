export interface Position {
  id: number
  ticker: string,
  price: number,
  quantity: number,
  'opened-at': string
}

export interface ClosedPosition extends Position {
  'closed-at': string,
  'close-price': number,
  profit: number,
}
