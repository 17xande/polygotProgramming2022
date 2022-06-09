const a: number[] = [];
const b = a;

b.push(1);

// I just changed a constant... and b is a pointer to a... in JavaScript. wat.
// console.log(a);

type fooo = {
  bar?: string;
}

function doSomething(f: fooo): boolean {
  if (f.bar) {
    return true;
  }

  return false;
}

enum TSEnum {
  Foo,
  Bar,
  Baz,
}

console.log(TSEnum);