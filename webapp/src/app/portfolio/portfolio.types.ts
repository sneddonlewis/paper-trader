import { ClosedPosition, Position } from '../position/position.model';

export interface Portfolio {
  id: number;
  userID: number;
  name: string;
  value: number; // TODO rename this amount or cash, value includes unrealised P&L
  openPositions: Position[];
  closedPositions: ClosedPosition[];
}
