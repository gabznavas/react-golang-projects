'use client'

import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "./ui/card";
import { Project } from "@/services/types";
import { useRouter } from "next/navigation";
import { Button } from "./ui/button";
import { PencilIcon, TrashIcon } from "lucide-react";

import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog"

import { useState } from "react";
import { requests } from "@/services/api";
import { toast } from "sonner";

type Props = {
  project: Project;
  onDelete: (id: string) => void;
}

export const ProjectCard = ({ project, onDelete }: Props) => {
  const [isOpenDialogToDelete, setIsOpenDialogToDelete] = useState(false);
  const router = useRouter();

  const handleOnUpdate = () => {
    router.push(`/project/details?id=${project.id}`);
  }

  const handleOnDelete = async () => {
    if (!isOpenDialogToDelete) return
    if (!project.id) return

    try {
      await requests.deleteProject(project.id);
      toast.success('Projeto deletado com sucesso', {
        position: "top-center",
        duration: 3000,
        style: {
          background: '#000',
          color: '#fff',
        },
      });
      onDelete(project.id);
    } catch (error) {
      console.error(error);
    } finally {
      setIsOpenDialogToDelete(false);
    }
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
        <Button onClick={() => setIsOpenDialogToDelete(true)}
          variant="outline"
          size="icon"
          className="bg-red-500 text-white cursor-pointer w-2/4">
          <TrashIcon />
        </Button>
        <Button
          onClick={() => handleOnUpdate()}
          variant="outline"
          size="icon"
          className="bg-yellow-500 text-white cursor-pointer w-2/4">
          <PencilIcon />
        </Button>

        <AlertDialog open={isOpenDialogToDelete} onOpenChange={setIsOpenDialogToDelete}>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Tem certeza que deseja remover o projeto?</AlertDialogTitle>
              <AlertDialogDescription>
                <span className="text-red-500 font-normal">
                  Esta ação não pode ser desfeita. Isso irá deletar o projeto permanentemente.
                </span>
              </AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel
                className="bg-gray-500 text-white cursor-pointer w-2/4 hover:bg-gray-600">
                Cancelar
              </AlertDialogCancel>
              <AlertDialogAction onClick={() => handleOnDelete()}
                className="bg-red-500 text-white cursor-pointer w-2/4 hover:bg-red-600">
                Remover
              </AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      </CardFooter>
    </Card >
  );
};