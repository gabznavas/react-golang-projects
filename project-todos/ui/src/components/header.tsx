import Link from "next/link";

export const Header = () => {
  return (
    <header className="flex bg-gray-100 justify-between items-center p-4">
      <h1 className="text-2xl font-bold">
        <Link href="/project/list">Projetos</Link>
      </h1>
    </header>
  );
};