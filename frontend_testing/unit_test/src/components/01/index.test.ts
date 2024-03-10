import { add, sub, RangeError } from "./";

describe("四則演算", () => {
  describe("add", () => {
    test("戻り値は、第一引数と第二引数の「和」である", () => {
      expect(add(50, 50)).toBe(100);
    });
    test("合計の上限は、'100'である", () => {
      expect(add(70, 80)).toBe(100);
    });
    test("引数が'0〜100'の範囲外だった場合、例外をスローする", () => {
      const message = "入力値は0〜100の間で入力してください";
      expect(() => add(-10, 10)).toThrow(message);
      expect(() => add(-10, 10)).toThrow(RangeError);
      expect(() => add(10, -10)).toThrow(message);
      expect(() => add(10, -10)).toThrow(RangeError);
      expect(() => add(-10, 110)).toThrow(message);
      expect(() => add(-10, 110)).toThrow(RangeError);
    });
  });

  describe("sub", () => {
    test("戻り値は、第一引数と第二引数の「差」である", () => {
      expect(sub(51, 50)).toBe(1);
    });
    test("戻り値の合計は、下限が'0'である", () => {
      expect(sub(70, 80)).toBe(0);
    });
  });
});
