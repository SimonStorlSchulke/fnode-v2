export interface NodeTree {
  Nodes: FNode[]
  Links: NodeLink[]
}

export interface FNode {
  Type: string
  Id: string
  Inputs: FInput[]
  Outputs: FOutput[],
  Options: any,
  Meta: {
    PosX: number,
    PosY: number
  }
}

export interface FOutput {
  Name: string,
  Type: string,
}

export interface FInput {
  Name: string,
  Type: string,
  DefaultValue: any,
}

export interface FOption {
  Choices:        string[]
  SelectedOption: string}

export interface Mode {
  SelectedOption: string
}

export interface NodeLink {
  FromNode: string
  FromOutput: number
  ToNode: string
  ToInput: number
}