import type { Sheet } from "@/types/proto-es/v1/sheet_service_pb";
import type { Worksheet } from "@/types/proto-es/v1/worksheet_service_pb";

export const extractSheetUID = (name: string) => {
  const pattern = /(?:^|\/)sheets\/([^/]+)(?:$|\/)/;
  const matches = name.match(pattern);
  return matches?.[1] ?? "-1";
};

export const isLocalSheet = (name: string) => {
  return extractSheetUID(name).startsWith("-");
};

export const setSheetStatement = (
  sheet: Sheet | Worksheet,
  statement: string
) => {
  sheet.content = new TextEncoder().encode(statement);
  sheet.contentSize = BigInt(new TextEncoder().encode(statement).length);
};

export const getSheetStatement = (sheet: Sheet | Worksheet) => {
  return new TextDecoder().decode(sheet.content);
};

export const extractSheetCommandByIndex = (sheet: Sheet, index: number) => {
  const commands = sheet.payload?.commands;
  if (!commands) return undefined;
  const command = commands[index];
  if (!command) return undefined;
  const subarray = sheet.content.subarray(command.start, command.end);
  return new TextDecoder().decode(subarray);
};
