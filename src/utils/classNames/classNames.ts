export type Modes = Record<string, string | boolean | undefined>

export function classNames(
  cls: string,
  additional: Array<string | undefined> = [],
  modes: Modes = {}
): string {
  return [
    cls,
    ...additional.filter(Boolean),
    ...Object.entries(modes)
      .filter(([, value]) => Boolean(value))
      .map(([classname]) => classname)
  ].join(' ')
}
