

export function ext(name: { match: (arg0: RegExp) => any[]; }) {
    return name.match(/\.([^.]+)$/)?.[1]
}