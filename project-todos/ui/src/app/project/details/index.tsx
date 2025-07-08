import { useParams } from "next/navigation";

export default function ProjectDetails() {
  const { id } = useParams();
  return <div>Project {id}</div>;
}