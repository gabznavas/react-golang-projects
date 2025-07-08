'use client'

import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";
import { Project } from "@/services/types";
import { useRouter } from "next/navigation";

type Props = {
  project: Project;
}

export const ProjectCard = ({ project }: Props) => {
  const router = useRouter();

  return (
    <Card
      className="min-w-[320px] h-[300px] cursor-pointer hover:shadow-lg transition-all duration-300"
      onClick={() => router.push(`/project/details?id=${project.id}`)}>
      <CardHeader>
        <CardTitle>{project.name}</CardTitle>
      </CardHeader>
      <CardContent>
        <p>{project.description}</p>
      </CardContent>
    </Card>
  );
};