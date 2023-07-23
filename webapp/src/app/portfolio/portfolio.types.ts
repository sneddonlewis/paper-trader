import { ClosedPosition, Position } from '../position/position.model';

export interface Portfolio {
  id: number;
  userID: number;
  name: string;
  value: number;
  openPositions: Position[];
  closedPositions: ClosedPosition[];
}
