package article_application

import (
	"context"

	"github.com/williamu04/medium-clone/domain/model"
	"github.com/williamu04/medium-clone/domain/repository"
)

type RetrieveAllArticleUseCase struct {
	repository repository.ArticleRepository
	topic      repository.TopicRepository
}

func NewRetrieveAllArticlesUseCase(repository repository.ArticleRepository, topic repository.TopicRepository) *RetrieveAllArticleUseCase {
	return &RetrieveAllArticleUseCase{repository: repository, topic: topic}
}

type RetrieveAllArticlesOutput struct {
	Articles []RetrieveArticleOutput
}

func (uc *RetrieveAllArticleUseCase) Execute(ctx context.Context, filter map[string]any) (*RetrieveAllArticlesOutput, error) {
	articles, err := uc.repository.FindAllArticles(ctx, filter)

	if err != nil {
		return nil, err
	}

	if articles == nil {
		return nil, nil
	}

	return &RetrieveAllArticlesOutput{
		Articles: mapArticlesToOutputs(articles),
	}, nil
}

func mapArticlesToOutputs(articles []*model.ArticleModel) []RetrieveArticleOutput {
	outputs := make([]RetrieveArticleOutput, len(articles))

	for i, article := range articles {
		outputs[i] = RetrieveArticleOutput{
			ID:          article.ID,
			Title:       article.Title,
			Description: article.Description,
			Body:        article.Body,
			AuthorID:    article.AuthorID,
		}
	}
	return outputs
}
