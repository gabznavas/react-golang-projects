'use client'

import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "./ui/card";
import { Project } from "@/services/types";
import { useRouter } from "next/navigation";
import { Button } from "./ui/button";
import { PencilIcon, TrashIcon } from "lucide-react";

type Props = {
  project: Project;
}

export const ProjectCard = ({ project }: Props) => {
  const router = useRouter();

  const handleOnUpdate = () => {
    router.push(`/project/details?id=${project.id}`);
  }

  const handleOnDelete = () => {
    console.log('delete');
  }

  return (
    <Card
      className="flex flex-col min-w-[320px] h-[300px]">
      <CardHeader className="flex-1">
        <CardTitle>{project.name}</CardTitle>
      </CardHeader>
      <CardContent className="flex-8">
        <p>{project.description}</p>
      </CardContent>
      <CardFooter className="flex gap-2 flex-1">
        <Button onClick={() => handleOnDelete()} variant="outline" size="icon" className="bg-red-500 text-white cursor-pointer w-2/4">
          <TrashIcon />
        </Button>
        <Button onClick={() => handleOnUpdate()} variant="outline" size="icon" className="bg-yellow-500 text-white cursor-pointer w-2/4">
          <PencilIcon />
        </Button>
      </CardFooter>
    </Card >
  );
};