// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {core} from '../models';
import {context} from '../models';

export function AddLink(arg1:core.NodeLink):Promise<void>;

export function AddNode(arg1:string,arg2:number,arg3:number):Promise<void>;

export function ClearTree():Promise<void>;

export function GetNodeCategories():Promise<Array<core.NodeCategory>>;

export function GetTestTree():Promise<core.SerializableTree>;

export function ParseTree():Promise<void>;

export function RemoveLink(arg1:core.NodeLink):Promise<void>;

export function SetContext(arg1:context.Context):Promise<void>;

export function UpdateInputDefaultValue(arg1:string,arg2:number,arg3:any,arg4:number):Promise<void>;

export function UpdateNodePosition(arg1:string,arg2:number,arg3:number):Promise<void>;

export function UpdateUption(arg1:string,arg2:string,arg3:string):Promise<boolean>;