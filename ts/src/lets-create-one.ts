export interface Set<T> {
  add (item: T): void

  contains (item: T): boolean

  equals (other: Set<T>): boolean
}

