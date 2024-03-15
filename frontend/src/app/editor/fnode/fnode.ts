export interface FTree {
  Nodes: FNode[]
  Links: NodeLink[]
}

export interface FNode {
  Type: string
  Id: string
  Inputs: FInput[]
  Outputs: FOutput[],
  Options: Map<string, NodeOption>,
  Meta: {
    PosX: number,
    PosY: number,
    Category: string,
  }
}

export interface NodeOption {
  Choices:        string[],
  SelectedOption: string,
}

export interface FOutput {
  Name: string,
  Type: number,
}

export const FType = {
  Float: 0,
  Int: 1,
  String: 2,
  Bool: 3,
  StringList: 4,
  FloatList: 5,
  IntList: 6,
};

export interface FInput {
  Name: string,
  Type: number,
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

export interface NodeCategory {
  Name:      string
  NodeTypes: string[]
}

export interface FileList {
  LooseFiles: string[];
  Directories: {
      Path: string;
      Recursive: boolean;
  }[];
}