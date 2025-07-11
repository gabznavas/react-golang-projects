'use client'

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { PlusIcon, SearchIcon } from "lucide-react";
import { ProjectCard } from "@/components/project-card";
import { useEffect, useState } from "react";
import { requests } from "@/services/api";
import { Project } from "@/services/types";
import { useRouter } from "next/navigation";
import { Form } from "@/components/ui/form";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import { AxiosError } from "axios";
import { toast } from "sonner";

const schema = yup.object().shape({
  search: yup.string(),
  offset: yup.number(),
  limit: yup.number(),
});

const defaultSearch = {
  limit: 10,
  offset: 0,
  search: '',
}

type FormSchema = yup.InferType<typeof schema>;

export default function ProjectList() {
  const [projects, setProjects] = useState<Project[]>([]);
  const router = useRouter();

  useEffect(() => {
    const fetchProjects = async () => {
      const projects = await requests.getProjects(defaultSearch);
      setProjects(prev => {
        const newProjects = projects.filter((project: Project) => !prev.some(p => p.id === project.id));
        return [...prev, ...newProjects];
      });
    }
    fetchProjects();
  }, []);

  const handleCreateProject = () => {
    router.push('/project/details');
  }

  const form = useForm({
    resolver: yupResolver(schema),
    defaultValues: defaultSearch,
  });

  const handleSearch = async (formData: FormSchema) => {
    try {
      const projects = await requests.getProjects({
        limit: formData.limit || defaultSearch.limit,
        offset: formData.offset || defaultSearch.offset,
        search: formData.search || defaultSearch.search,
      });
      setProjects(projects);
    } catch (error) {
      if (error instanceof AxiosError && error.response?.data.error) {
        toast.error(error.response.data.error, {
          position: "top-center",
        });
      }
      toast.error("Erro ao buscar projetos", {
        position: "top-center",
      });
    }
  }

  const handleDeleteProject = (id: string) => {
    setProjects(prev => prev.filter(project => project.id !== id));
  }

  return (
    <div className="flex flex-col gap-4 mx-auto">
      <Form {...form}>
        <form className="flex gap-4 justify-between items-center" onSubmit={form.handleSubmit(handleSearch)}>
          <Input type='search' placeholder="Procure por um projeto" {...form.register('search')} />
          <Button autoFocus variant="outline" className="flex items-center gap-2 w-1/6 cursor-pointer">
            <SearchIcon />
          </Button>
          <Button variant="default" className="w-1/10 cursor-pointer" type="button" onClick={handleCreateProject}>
            <PlusIcon />
          </Button>
        </form>
      </Form>
      <div className="grid grid-cols-1 md:grid-cols-2 2xl:grid-cols-4 gap-4 w-full">
        {
          projects.map((project, index) => (
            <ProjectCard key={index} project={project} onDelete={handleDeleteProject} />
          ))
        }
      </div>
      <div className="flex gap-1 justify-center items-center ">
        {
          new Array(10).fill(0).map((_, index) => (
            <div key={index} className="flex items-center justify-center text-sm text-gray-600 w-10 h-10 bg-gray-200 animate-pulse rounded-md">
              {index + 1}
            </div>
          ))
        }
      </div>
    </div>
  )
}