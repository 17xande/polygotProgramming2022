import path from 'path'
import { Opts } from './opts'

export enum Operation {
  Print,
  Add,
  Remove,
}

export type Config = {
  args: string[],
  operation: Operation,
  config: string,
  pwd: string,
}

export default function getConfig(opts: Opts): Config {
  return {
    pwd: getPwd(opts),
    config: config(opts),
    args: getArgs(opts),
    operation: getOperation(opts),
  }
}

function getPwd(opts: Opts): string {
  if (opts.pwd) {
    return opts.pwd
  }

  return process.cwd()
}

function config(opts: Opts): string {
  if (opts.config) {
    return opts.config
  }

  const home = process.env["HOME"]
  const loc = process.env["XDG_CONFIG_HOME"] || home
  if (!loc) {
    throw new Error("unable to determine config location");
  }

  if (loc === home) {
    return path.join(loc, ".projector.json")
  }

  return path.join(loc, "projector", "projector.json")
}

function getArgs(opts: Opts): string[] {
  if (!opts.args || opts.args.length === 0) {
    return []
  }
  
  const operation = getOperation(opts)
  if (operation === Operation.Print) {
    if (opts.args.length > 1) {
      throw new Error(`expected 0 or 1 arguments but got ${opts.args.length}`)
    }
    return opts.args
  }

  if (operation === Operation.Add) {
    if (opts.args.length !== 3) {
      throw new Error(`expected 2 but got ${opts.args.length}`)
    }
    return opts.args.slice(1)
  }

  if (opts.args.length !== 1) {
    throw new Error(`expected 1 but got ${opts.args.length}`)
  }
  return opts.args.slice(1)

}

function getOperation(opts: Opts): Operation {
  if (!opts.args || opts.args.length === 0) {
    return Operation.Print;
  }

  if (opts[0] === "add") {
    return Operation.Add
  }

  if (opts[0] === "rm") {
    return Operation.Remove
  }

  return Operation.Print
}
