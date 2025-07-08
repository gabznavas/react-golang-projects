import { Header } from "@/components/header";
import ProjectList from "./index";

export default function ProjectListPage() {
  return (
    <div className="flex flex-col gap-4">
      <div className="flex flex-col gap-4">
        <Header />
        <ProjectList />
      </div>
    </div>
  );
}